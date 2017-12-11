package fetcher

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"testing"
	"time"
)

const tempDest = "testdata/tmp.txt"
const baseTestURL = "http://localhost:7777"

// TestMain kicks off a simple server that allows us to test the
// Fetch method of the BasicRSFetcher and prove we handle errors appropriately
func TestMain(m *testing.M) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/200", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not Found, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/403", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Forbidden, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/503", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "Gateway timeout, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":7777", nil)
	os.Exit(m.Run())
}

func TestStatusOKWritesToDisk(t *testing.T) {
	type testData struct {
		path        string
		expResponse int
		expStatus   int
		expErr      error
	}

	testTable := []testData{
		testData{
			path:      "/200",
			expStatus: http.StatusOK,
			expErr:    nil,
		},
		testData{
			path:      "/404",
			expStatus: http.StatusNotFound,
			expErr:    ErrNon200Response,
		},
		testData{
			path:      "/403",
			expStatus: http.StatusForbidden,
			expErr:    ErrNon200Response,
		},
		testData{
			path:      "/503",
			expStatus: http.StatusBadGateway,
			expErr:    ErrNon200Response,
		},
	}
	brf := &BasicRSFetcher{}
	for _, td := range testTable {
		status, err := brf.Fetch(fmt.Sprintf("%s%s", baseTestURL, td.path), tempDest)
		if err != td.expErr {
			t.Errorf("Unexpected error returned: %v Exp: %v", err, td.expErr)
		}
		if status != td.expStatus {
			t.Errorf("Unexpected status code: %d; Exp: %d", status, td.expStatus)
		}
	}
}
