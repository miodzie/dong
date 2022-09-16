package domain

type Dong struct {
	Emoji    string
	Category string
	Author   string
}

type RandomDonger interface {
	Dong() (Dong, error)
	DongByCategory(string) (Dong, error)
}

type DongRepository interface {
	Count() (int64, error)
	CountWithCategory(string) (int64, error)
	Categories() ([]string, error)
	ByCategory(string) ([]Dong, error)
	ById(int64)
}

func (d Dong) String() string {
	return d.Emoji
}
