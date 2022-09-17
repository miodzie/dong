package dong

type Emoji struct {
	Text     string
	Category string
	Author   string
}

type Repository interface {
	Random() (Emoji, error)
	RandomByCategory(string) (Emoji, error)
	Count() int64
	Categories() ([]string, error)
}

func (d Emoji) String() string {
	return d.Text
}
