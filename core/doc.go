// package core holds the main data structures and methods for reading CORE article metadata JSON.
// The main public functions allow you to either send in a URL ad apiKey - Process(). Or send in
// []byte from an article page JSON - ExtractArticle(). In both cases you get a ArticleWrapper object back with the
// extracted data available for further inspection and use.

package core
