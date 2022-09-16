package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/miodzie/dong/impl"
	"github.com/miodzie/dong/interactors"
	"os"
	"os/user"
	"path"
)

var workDir string

func main() {
	createWorkDir()
	controller := interactors.NewRandomDongInteractor(initDatabase())
	resp := controller.Handle(interactors.RandomDongReq{})
	if resp.Error != nil {
		fmt.Println("Error: ", resp.Error)
	}
	fmt.Println(resp.Emoji)
}

func createWorkDir() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	workDir = path.Join(usr.HomeDir, ".dong")

	if _, err := os.Stat(workDir); os.IsNotExist(err) {
		err := os.Mkdir(workDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func initDatabase() *impl.GormRepository {
	db, err := gorm.Open("sqlite3", path.Join(workDir, "dongs.db"))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return impl.NewGormRepository(db)
}
