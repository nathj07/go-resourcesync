package resourcesync

import "github.com/nathj07/go-resourcesync/fetcher"

// ResourceSync is the top level structure needed to interact with ResourceSync endpoints
type ResourceSync struct {
	Fetcher fetcher.RSFetcher
}

// New is the simplest way to instantiate a ready to use ResourceSync object
func New(f fetcher.RSFetcher) *ResourceSync {
	return &ResourceSync{
		Fetcher: f,
	}
}

// TODO:
// Function to retrieve data
// Function to parse data - this should be exported and take a []byte to parse returning an exported struct
// Function to do it all in one go?
// I think that for now the client will be very simple and it will be up to the caller to
// determine how far to follow the chain.
// In the future this could be changed with a slight variation/addition to the public API
