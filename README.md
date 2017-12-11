# go-resourcesync
A Go client for the ResourceSynch protocol.

## Useful Links and Further Reading
[ResourceSync Specification](http://www.openarchives.org/rs/toc)

## Code Structure
`resourcesync` This package holds the main data structures use in working with the resourcesync protocol

`fetcher` describes a simple iterface for HTPP fetching and contais `basicFetcher` a simplistic implemntation used in the CLI tool.

`cmd` holds the CLI tool which can be useful in testing endpoints ahead of using them in your production application. This may also be helpful in debugging any issues as it enables you to see the actual response.
