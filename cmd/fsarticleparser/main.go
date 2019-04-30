package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/nathj07/go-resourcesync/core"
	"io/ioutil"
	"os"
)

var jsonFile = flag.String("file", "", "--file full path to a JSON file conforming to FastSync article schema")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of fsarticletester:
fsarticletester is a tool within the go-resourcesync client that allows you parse CORE FastSync JSON and view the Go struct output  
Flags:`)
		flag.PrintDefaults()
	}
	flag.Parse()
	if *jsonFile == "" {
		fmt.Fprintf(os.Stderr, "--file is a mandatory argument, please specify an absolute path to the JSON file \n")
		os.Exit(1)
	}
	readJsonFile()
}

func readJsonFile(){
	data, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file: %q - %v\n", *jsonFile, err)
		os.Exit(2)
	}
	res := &core.FSArticle{}
	err = json.Unmarshal(data, res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse JSON: %v\n", err)
		os.Exit(3)
	}
	spew.Dump(res)
}
