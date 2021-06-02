# GoLearnGo

## Resources

https://dave.cheney.net/resources-for-new-go-programmers

I think this is a pretty good enough curation of all the resources one will
ever need to "Get Started" with Go.

## Progress

1. I have started with https://tour.golang.org
2. I am making a [client](
   https://github.com/mbtamuli/GoLearnGo/blob/master/doclient/doclient.go)
   to consume the [DigitalOcean API](
   https://developers.digitalocean.com/documentation/v2/) to understand the
   concepts better. Looking at the source of [godo](
   https://github.com/digitalocean/godo) as inspiration.
3. Done with [go-koans](https://github.com/mbtamuli/go-koans/)
4. Experimenting with [GitHub v4 GraphQL API](githubclient/)
5. Gophercises quiz
6. Random snippets added

## Notes
If using `database/sql`, you can log sql statements using

```go
import (
	"github.com/luna-duclos/instrumentedsql"
	"modernc.org/sqlite"
)

logger := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
	log.Printf("%s %v", msg, keyvals)
})

sql.Register("instrumented-sqlite", instrumentedsql.WrapDriver(&sqlite3.SQLiteDriver{}, instrumentedsql.WithLogger(logger)))
db, err := sql.Open("instrumented-sqlite", dataSourceName)
```
