package impl

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/miodzie/dong"
	"path"
	"testing"
)

func TestGormRepository_Save(t *testing.T) {
	db, err := gorm.Open("sqlite3", path.Join("", "dongs.sqlite"))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	repository := NewGormRepository(db)

	err = repository.Save([]dong.Emoji{{Text: ":D", Category: "happy"}})
	if err != nil {
		panic(err)
	}
}
