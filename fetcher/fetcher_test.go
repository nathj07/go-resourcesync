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
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Not Found, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/403", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Forbidden, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/503", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Millisecond)
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Gateway timeout, %q", html.EscapeString(r.URL.Path))
	})

	go http.ListenAndServe(":7777", nil)
	os.Exit(m.Run())
}

func TestStatusOKWriutesToDiks(t *testing.T) {
	brf := &BasicRSFetcher{}
	err := brf.Fetch(fmt.Sprintf("%s%s", baseTestURL, "/200"), tempDest)
	if err != nil {
		t.Fatalf("Unexpected error fetching data: %v", err)
	}
}
