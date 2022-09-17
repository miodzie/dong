package impl

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/miodzie/dong"
	"path"
	"testing"
)

// Just a crappy smoke test i'm a bum
func TestGormRepository_Save(t *testing.T) {
	db, err := gorm.Open("sqlite3", path.Join("", "dongs.sqlite"))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.AutoMigrate(Dong{})
	repository := NewGormRepository(db)
	var emojis []dong.Emoji
	emojis = append(emojis, dong.Emoji{Text: ":D", Category: "happy"})
	fmt.Printf("%+v\n", emojis)

	err = repository.Save(emojis)
	if err != nil {
		panic(err)
	}
}
