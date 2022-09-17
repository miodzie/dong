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
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)

	}
}

func run(args []string) error {
	createWorkDir()
	controller := interactors.NewRandomDongInteractor(initDatabase())
	req := interactors.RandomDongReq{}
	if len(args) > 1 {
		req.Category = args[1]
	}
	resp := controller.Handle(req)
	if resp.Error != nil {
		return resp.Error
	}
	fmt.Println(resp.Emoji)
	return nil
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
