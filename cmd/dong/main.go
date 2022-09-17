package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/miodzie/dong"
	"github.com/miodzie/dong/impl"
	"github.com/miodzie/dong/interactors"
	"os"
	"os/user"
	"path"
)

var workDir string
var repository dong.Repository

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	createWorkDir()
	initDatabase()

	handleCommands(args)

	return printRandomDong(args)
}

func printRandomDong(args []string) error {
	req := interactors.RandomDongReq{}
	if len(args) > 0 {
		req.Category = args[0]
	}
	controller := interactors.NewRandomDongInteractor(repository)
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

func initDatabase() {
	db, err := gorm.Open("sqlite3", path.Join(workDir, "dongs.db"))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	repository = impl.NewGormRepository(db)
}
