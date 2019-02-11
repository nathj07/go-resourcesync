package main

import (
	"flag"
	"fmt"
	"github.com/nathj07/go-resourcesync/core"
	"log"
	"net/url"
	"os"

	"github.com/nathj07/go-resourcesync/fetcher"
	"github.com/nathj07/go-resourcesync/resourcesync"
)

var (
	target     = flag.String("target", "", "--target=http:/example.com/resourcesync.xml")
	targetType = flag.String("targettype", "resourcesync", "--targettype indicates if this is 'resourecesync' or 'core'")
	follow     = flag.Bool("follow", false, "--follow indicates if resource link index sets should be followed until resource lists are reached")
	apiKey     = flag.String("apikey", "", "--apikey is used in requests for CORE article metadata")
	verbose    = flag.Bool("verbose", false, "--verbose, if set will print all the links discovered")
)

const (
	segmentation = "====================" // breaks up the output
	targetCore   = "core"
)

type app struct {
	rs             *resourcesync.ResourceSync
	ce             *core.Extractor
	follow         bool
	verbose        bool
	indexLinkCount int
	linkCount      int
	startingPoint  string
	targetType     string
	apiKey         string
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
	if *targetType == targetCore && *apiKey == "" {
		log.Printf("When specifying the target type %q you must provide an API Key for use in requests for the metadata.\n", targetCore)
		os.Exit(1)
	}
	app := &app{
		rs: &resourcesync.ResourceSync{
			Fetcher: &fetcher.BasicRSFetcher{},
		},
		ce: &core.Extractor{
			Fetcher: &fetcher.BasicRSFetcher{},
		},
		follow:        *follow,
		verbose:       *verbose,
		startingPoint: *target,
		targetType:    *targetType,
		apiKey:        *apiKey,
	}

	app.processTarget(*target)
	if app.verbose {
		log.Println("Starting point:", app.startingPoint)
		log.Println("Index links found:", app.indexLinkCount)
		log.Println("Resource links found:", app.linkCount)
	}
}

func (app *app) processTarget(target string) {
	if app.targetType == targetCore {
		// fetch and unmarshall the data
		app.processCoreMetadata()
		return
	}
	app.processResourceSync(target)
}

func (app *app) processCoreMetadata() {
	data, err := app.ce.Process(app.startingPoint, app.apiKey)
	if err != nil {
		log.Printf("failed to process CORE metadata from %q: %v\n", app.startingPoint, err)
		os.Exit(1)
	}
	log.Println(data)
}

func (app *app) processResourceSync(target string) {
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
				app.processResourceSync(index.Loc)
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
