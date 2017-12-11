package fetcher

// Fetcher is a small interface definning the main act of fetching resources.
// This can be overwritten by the user of the client to provide more custom fetch behaviour
type Fetcher interface {
	Fetch(source, dest string) error
}
