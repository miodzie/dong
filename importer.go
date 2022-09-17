package dong

// Fetcher fetches dongs in any fashion desired.
// There is a default Web Scraper implementation of this from
// http://dongerlist.com in the commandline package.
// It returns the dongs found and any error encountered.
type Fetcher interface {
	Fetch() ([]Emoji, error)
}
