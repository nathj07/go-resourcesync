package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Fetcher is a small interface definning the main act of fetching resources.
// This can be overwritten by the user of the client to provide more custom fetch behavior
type Fetcher interface {
	Fetch(source, dest string) error
}

// BasicRSFetcher is a simple implementation of the Fetcher interface. It is safe to use
// but limited in capability. No timeouts or extra headers are defined. The general recommendation
// is that the user of this client write their own implementation of the Fetcher interface.
type BasicRSFetcher struct{}

// Fetch retrieves the resource from source and writes it to dest. It is the callers responsibility
// to clear up any local files when they are finished with.
// This fetcher implementation will return an error for a non-200 response.
func (brf *BasicRSFetcher) Fetch(source, dest string) error {
	res, err := http.Get(source)
	if err != nil {
		return fmt.Errorf("error making GET request against: %q: %v", source, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("non-%d status code %d returned. Data not written", http.StatusOK, res.StatusCode)
	}
	brf.writeToDisk(dest, res.Body)
	if err != nil {
		return err
	}

	return nil
}

// writeToDisk does just that, writing the contents of the supplied io.Reader to the stated destination.
func (brf *BasicRSFetcher) writeToDisk(dest string, content io.Reader) error {
	contentFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("error creating local file for download: %v", err)
	}
	defer contentFile.Close()
	_, err = io.Copy(contentFile, content)
	if err != nil {
		return fmt.Errorf("error copying response body to local file: %v", err)
	}
	return nil
}
