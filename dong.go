package dong

// Fetcher fetches dongs in any fashion desired.
// There is a default Web Scraper implementation of this from
// http://dongerlist.com in the commandline package.
// It returns the dongs found and any error encountered.
type Fetcher interface {
	Fetch() ([]Emoji, error)
}

type Repository interface {
	Random() (Emoji, error)
	RandomByCategory(string) (Emoji, error)
	Count() int64
	Categories() ([]string, error)
	Save([]Emoji) error
}

type Emoji struct {
	Text     string
	Category string
}

func (d Emoji) String() string {
	return d.Text
}
