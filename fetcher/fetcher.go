package fetcher

// Fetcher is a small interface definning the main act of fetching resources.
// This can be overwritten by the user of the client to provide more custom fetch behaviour
type Fetcher interface {
	Fetch(source, dest string) error
}

// BasicRSFetcher is a simple implementation of the Fetcher interface. It is safe to use
// but limited in capability. No timeouts or extra headers are defined. The general recommendation
// is that the user of this client write their own implementation of the Fetcher interface.
type BasicRSFetcher struct{}

// Fetch retrieves the resource from source and writes it to dest. It is the callers responsibility
// to clear up any local files when they are finished with.
func (brf *BasicRSFetcher) Fetch(source, dest string) error {

}
