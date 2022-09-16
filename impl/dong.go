package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/miodzie/dong/domain"
)

type Dong struct {
	gorm.Model
	Dong     string
	Category string
}

func (d Dong) ToDomainDong() domain.Dong {
	return domain.Dong{
		Emoji:    d.Dong,
		Category: d.Category,
	}
}

func (d Dong) String() string {
	return d.Dong
}
