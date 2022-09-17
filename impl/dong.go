package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/miodzie/dong"
)

// Dong TODO: Migrate column names?
type Dong struct {
	gorm.Model
	Dong     string
	Category string
}

func (d Dong) ToDomainDong() dong.Emoji {
	return dong.Emoji{
		Text:     d.Dong,
		Category: d.Category,
	}
}

func (d Dong) String() string {
	return d.Dong
}
