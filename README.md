# go-resourcesync
A Go client for the ResourceSync protocol.

This client makes no decisions on what to do with the data that is returned. It will simply fetch it, unmarshal it into a easy to use data structure and return that. It is up to the consumer to decide what to do with this data.

For example if processing at the index level it is up to the caller to decide to follow the links or not. If operating at the list level it is up to the caller to decide if they wish to fetch the PDF or the metadata. The purpose of this library is to simplify the fetching and unmarshaling.

## Useful Links and Further Reading
[ResourceSync Specification](http://www.openarchives.org/rs/toc)

## Code Structure
`resourcesync` This package holds the main data structures use in working with the resourcesync protocol

`fetcher` describes a simple interface for HTTP fetching and contains `basicFetcher` a simplistic implementation used in the CLI tool.

`cmd` holds the CLI tool which can be useful in testing endpoints ahead of using them in your production application. This may also be helpful in debugging any issues as it enables you to see the actual response.

Example:

```bash
cd cmd
go run main.go -target http://publisher-connector.core.ac.uk/resourcesync/sitemaps/Frontiers/metadata/resourcelist_0001.
xml -verbose
```
This command will fetch the specified target, which is the bottom of ResourceSync hierarchy, and print all the contain link information to stdout. This code alos serves as a useful example for using this library in your own code.
