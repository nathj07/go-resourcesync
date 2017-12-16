package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/nathj07/go-resourcesync/fetcher"
	"github.com/nathj07/go-resourcesync/resourcesync"
)

var (
	target      = flag.String("target", "", "--target=http:/example.com/resourcesync.xml")
	followIndex = flag.Bool("follow-index", false, "--follow-index means the tool will follow the links in a sitemaps index. Only one recursion is made")
	verbose     = flag.Bool("verbose", false, "--verbose, if set will print all the links discovered")
)

type app struct {
	rs          *resourcesync.ResourceSync
	endpoint    string
	followIndex bool
	verbose     bool
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:
%s is a tool within the go-resourcesync client that allows you to visit a resourcesync endpoint and evaluate the response.
Flags:
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	u, err := url.Parse(*target)
	if err != nil || !u.IsAbs() {
		log.Println("This tool expects a --target flag to be passed in with a valid, absolute URL.")
		os.Exit(1)
	}

	app := &app{
		rs: &resourcesync.ResourceSync{
			Fetcher: &fetcher.BasicRSFetcher{},
		},
		endpoint:    *target,
		followIndex: *followIndex,
		verbose:     *verbose,
	}
	resources, err := app.checkResourceSync()
	if err != nil {
		log.Printf("Error encountered checking resourcesync: %v\n", err)
		os.Exit(2)
	}

	if !app.followIndex {
		log.Println("ResourceSync Data:")
		app.printLinks(resources.RL.URLSet)
		os.Exit(0)
	}
}
func (app *app) checkResourceSync() (*resourcesync.ResourceData, error) {
	// fetch it
	data, status, err := app.rs.Fetcher.Fetch(app.endpoint)
	if err != nil {
		return nil, err
	}
	log.Printf("Return status code: %d\n", status)

	return app.rs.Parse(data)
}

func (app *app) printLinks(resources []resourcesync.ResourceURL) {
	if app.verbose {
		for _, r := range resources {
			log.Printf("Loc: %+v\nLastMod: %+v\nrsmd: %+v\nrsln: %+v\n\n", r.Loc, r.LastMod, r.RSMD, r.RSLN)
		}
	}
	log.Printf("Total links found: %d\n", len(resources))
}
