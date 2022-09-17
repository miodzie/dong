package main

import (
	"fmt"
	"github.com/miodzie/dong/impl"
	"github.com/miodzie/dong/interactors"
	"os"
	"strings"
)

var commands map[string]func()

func init() {
	commands = make(map[string]func())
	commands["count"] = func() {
		fmt.Println(repository.Count())
		os.Exit(0)
	}
	// categories
	commands["cat"] = func() {
		cats, err := repository.Categories()
		if err != nil {
			panic(err)
		}
		fmt.Println(strings.Join(cats, ", "))
		os.Exit(0)
	}
	commands["scrape"] = func() {
		exec := interactors.NewScrapeDongsInteractor(impl.NewScraper(), repository)
		resp := exec.Handle()
		if resp.Error != nil {
			panic(resp.Error)
		}
		fmt.Println(resp.Message)
	}
	commands["version"] = func() {
		fmt.Println("ヽ༼ຈل͜ຈ༽ﾉ FOREVER DONG ヽ༼ຈل͜ຈ༽ﾉ")
		os.Exit(0)
	}
}

func handleCommands(args []string) {
	if len(args) == 0 {
		return
	}
	cmd, ok := commands[args[0]]
	if !ok {
		return
	}
	cmd()
}
