# go-resourcesync

[![CircleCI](https://circleci.com/gh/nathj07/go-resourcesync/tree/master.svg?style=svg&circle-token=0c9fb37da87f3dd9f9758a5c6fb279b626f760db)](https://circleci.com/gh/nathj07/go-resourcesync/tree/master)

A Go client for the ResourceSync protocol.

This client makes no decisions on what to do with the data that is returned. It will simply fetch it, unmarshal it into a easy to use data structure and return that. It is up to the consumer to decide what to do with this data.

For example if processing at the index level it is up to the caller to decide to follow the links or not. If operating at the list level it is up to the caller to decide if they wish to fetch the PDF or the metadata. The purpose of this library is to simplify the fetching and unmarshaling.

## Useful Links and Further Reading

[ResourceSync Specification](http://www.openarchives.org/rs/toc)

## Code Structure

`resourcesync` This package holds the main data structures use in working with the resourcesync protocol

`fetcher` describes a simple interface for HTTP fetching and contains `basicFetcher` a simplistic implementation used in the CLI tool.

`cmd` holds the CLI tool which can be useful in testing endpoints ahead of using them in your production application. This may also be helpful in debugging any issues as it enables you to see the actual response.

## Example CLI Usage

```bash
cd cmd
go run main.go -target http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/metadata/resourcelist_0001.
xml -verbose
```

This command will fetch the specified target, which is the bottom of ResourceSync hierarchy, and print all the contain link information to stdout. This code also serves as a useful example for using this library in your own code.

## Example Library Usage

```go
package main

import (
    "github.com/nathj07/go-resourcesync/fetcher"
    "github.com/ntahj07/go-resourcesync/resourcesync"
    "github.com/davecgh/go-spew/spew"
)

func main() {
    rs := &ResourceSync{
        Fetcher: fetcher.BasicFetcher{} // or your own fetcher conforming to the interface
    }
    // Using process
    data, err := rs.Process("http://resourcesync.org/endpoint")
    if err != nil {
        panic(err)
    }
    spew.Dump(data) // actually do something with it

    // Just using parse, having fetched the data to []byte
    parsedData, err := rs.Parse(rawData)
      if err != nil {
        panic(err)
    }
    spew.Dump(parsedData) // actually do something with it
}
```

The code in `cmd/main.go` also functions as a useful example of using this library.

## Contributing

Contributions are welcome and hopefully not to onerous to make. Please follow these simple steps:

1. Raise an issue; doesn't have to be a lot but something to connect the change to
1. Make a branch referencing the issue number
1. Open a PR and await feedback; please ensure you are up to date with master prior to opening the PR

In all code please follow the [standard Go idioms](https://golang.org/doc/effective_go.html); this just makes it easier for all users and maintainers of the code.

Alternatively if you wish to simply make an enhanceent request then simply raise an issue and leave it at that.

There's nothing more to it
