// pacakge goresourcesync is a simple go library for intertacting with resourcesync data.
// The resourcesync protocol is based on the pre-existing sitemap protocol and details can be found:
// http://www.openarchives.org/rs/toc
//
// The motivation for this protocol is easing the discovery of academic research content and providing a
// meachanism for updating those following the feed.
//
// With regards to this small library the miotivation is simply making is easy to consume the resourcesync data
// within a Go app.
//
// The library makes no decisions about waht to do with content, it does not automatically follow any links within
// the feed. Those decisions are left up to the caller.
//
// While the library does provide a simnple Fetcher implementation this is by no measn a production ready HTTP Fetcher.
// It is expected that the user of the library will implement their own fetcher either injecting it into the
// ResourceSync object or by fetching the data upfront first and simply calling the Parse function.

package goresourcesync
