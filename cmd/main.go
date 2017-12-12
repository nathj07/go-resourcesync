package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/nathj07/go-resourcesync/resourcesync"
)

var (
	target      = flag.String("target", "", "--target=http:/example.com/resourcesync.xml")
	followIndex = flag.Bool("follow-index", false, "--follow-index means the tool will follow the links in a sitemaps index. Only one recursion is made")
	verbose     = flag.Bool("verbose", false, "--verbose, if set will print all the links discovered")
)

// TODO switch fmt.Print to log.Info
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
		fmt.Println("This tool expects a --target flag to be passed in with a valid, absolute URL.")
		os.Exit(1)
	}

	resources, err := checkResourceSync(*target)
	if err != nil {
		fmt.Printf("Error encountered checking resourcesync: %v\n", err)
		os.Exit(2)
	}

	if !*followIndex {
		fmt.Println("ResourceSync Data:")
		printLinks(resources.URLSet)
		os.Exit(0)
	}
}

func checkResourceSync(addr string) (*resourcesync.ResourceList, error) {
	// fetch it
	b, err := retrieve(addr)
	if err != nil {
		return nil, err
	}

	// send it to parse
	rs := &resourcesync.ResourceSync{}
	return rs.Parse(b)
}

func retrieve(addr string) ([]byte, error) {
	res, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Returned Content Type: %s\n", res.Header.Get("Content-Type"))

	return b, nil
}

func printLinks(resources []resourcesync.ResourceURL) {
	if *verbose {
		for _, r := range resources {
			fmt.Printf("Loc: %+v\nLastMod: %+v\n\n", r.Loc, r.LastMod)
		}
	}
	fmt.Printf("Total links found: %d\n", len(resources))
}
