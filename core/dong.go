package core

import "github.com/jinzhu/gorm"

type Dong struct {
	gorm.Model
	Dong     string
	Category string
}

func (d Dong) String() string {
	return d.Dong
}
