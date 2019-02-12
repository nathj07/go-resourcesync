// package resourcesync holds the main data structures and methods for reading a ResourceSync feed.
// The main public functions allow you to either send in a ResourceSync feed URL - Process(). Or send in
// []byte from a ResourceSync feed - Parse(). In both cases you get a ResourceData object back with the
// parsed data available for further inspection and use.

package resourcesync
