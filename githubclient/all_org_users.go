package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type member struct {
	Login string
}

func main() {
	org := flag.String("org", "rtCamp", "Name of a GitHub Organization")
	flag.Parse()

	members := fetchUsers(context.Background(), *org, 100)

	for i := 0; i < len(members); i++ {
		fmt.Println(members[i].Login)
	}
}

func fetchUsers(ctx context.Context, orgName string, numberOfMembers int) []member {

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	var query struct {
		RateLimit struct {
			Cost      int
			Remaining int
			ResetAt   time.Time
		}
		Organization struct {
			Members struct {
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
				Nodes []member
			} `graphql:"members(first: $numberOfMembers, after: $membersCursor)"`
		} `graphql:"organization(login: $orgName)"`
	}

	variables := map[string]interface{}{
		"orgName":         githubv4.String(orgName),
		"numberOfMembers": githubv4.Int(numberOfMembers),
		"membersCursor":   (*githubv4.String)(nil),
	}

	var allMembers []member
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			panic(err)
		}
		allMembers = append(allMembers, query.Organization.Members.Nodes...)
		if !query.Organization.Members.PageInfo.HasNextPage {
			break
		}
		variables["membersCursor"] = githubv4.NewString(query.Organization.Members.PageInfo.EndCursor)
	}

	return allMembers
}
