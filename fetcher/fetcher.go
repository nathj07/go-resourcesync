package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RSFetcher is a small interface defining the main act of fetching resources.
// This can be overwritten by the user of the client to provide more custom fetch behavior
type RSFetcher interface {
	Fetch(source string) ([]byte, int, error)
}

// ErrNon200Response is returned from the BasicRSFetcher for any non-200 response
var ErrNon200Response = errors.New("non-200 status code returned")

// BasicRSFetcher is a simple implementation of the Fetcher interface. It is safe to use
// but limited in capability. No timeouts or extra headers are defined. The general recommendation
// is that the user of this client write their own implementation of the Fetcher interface.
type BasicRSFetcher struct{}

// Fetch retrieves the resource from source and writes it to dest. It is the callers responsibility
// to clear up any local files when they are finished with.
// This fetcher implementation will return an error for a non-200 response.
func (brf *BasicRSFetcher) Fetch(source string) ([]byte, int, error) {
	res, err := http.Get(source)
	if err != nil {
		return nil, 0, fmt.Errorf("error making GET request against: %q: %v", source, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, ErrNon200Response
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	return data, res.StatusCode, nil
}
