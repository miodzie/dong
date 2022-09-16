package domain

type Dong struct {
	Emoji    string
	Category string
	Author   string
}

type Repository interface {
	Random() (Dong, error)
	RandomByCategory(string) (Dong, error)
}

func (d Dong) String() string {
	return d.Emoji
}
