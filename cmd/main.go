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
	target  = flag.String("target", "", "--target=http:/example.com/resourcesync.xml")
	follow  = flag.Bool("follow", false, "--follow indicates if resource link index sets shoudl be followed until resource lists are reached")
	verbose = flag.Bool("verbose", false, "--verbose, if set will print all the links discovered")
)

var segmentation = "====================" // breaks up the output

type app struct {
	rs             *resourcesync.ResourceSync
	follow         bool
	verbose        bool
	indexLinkCount int
	linkCount      int
	startingPoint  string
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
		follow:        *follow,
		verbose:       *verbose,
		startingPoint: *target,
	}

	app.processTarget(*target)
	if app.verbose {
		log.Println("Starting point:", app.startingPoint)
		log.Println("Index links found:", app.indexLinkCount)
		log.Println("Resource links found:", app.linkCount)
	}
}

func (app *app) processTarget(target string) {
	resources, err := app.checkResourceSync(target)
	if err != nil {
		log.Printf("Error encountered checking resourcesync: %v\n", err)
		os.Exit(1)
	}

	log.Println("ResourceSync Data:")
	if resources.RL != nil {
		// this is the end of the chain, these we print, these would be the links to full text content
		if app.verbose {
			log.Println(resources.RL)
		}
		app.linkCount += len(resources.RL.URLSet)
	}
	if resources.RLI != nil {
		for _, index := range resources.RLI.IndexSet {
			if app.verbose {
				log.Println("Will follow:", index)
				log.Println(segmentation)
			}
			if app.follow {
				app.processTarget(index.Loc)
			}
		}
		if app.verbose {
			log.Printf("Followed %d index links\n", len(resources.RLI.IndexSet))
			log.Println(segmentation)
		}
		app.indexLinkCount += len(resources.RLI.IndexSet)
	}
}
func (app *app) checkResourceSync(target string) (*resourcesync.ResourceData, error) {
	// fetch it
	return app.rs.Process(target)
}
