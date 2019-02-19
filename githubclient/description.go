package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var logger = log.New(os.Stdout, ">>> ", log.Lshortfile)

func main() {
	flag.Parse()

	description, err := fetchRepoDescription(context.Background(), "mbtamuli", "playground")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(description)
}

func fetchRepoDescription(ctx context.Context, owner, name string) (string, error) {

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	var q struct {
		Repository struct {
			Description string
		} `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}

	err := client.Query(ctx, &q, variables)
	return q.Repository.Description, err
}
