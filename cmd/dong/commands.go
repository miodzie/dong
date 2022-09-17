package main

import (
	"fmt"
	"github.com/miodzie/dong/impl"
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
		scraper := &impl.Scraper{Domain: "http://dongerlist.com"}
		err := scraper.Run()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
	commands["version"] = func() {
		fmt.Println("ヽ༼ຈل͜ຈ༽ﾉ FOREVER DONG ヽ༼ຈل͜ຈ༽ﾉ")
		os.Exit(0)
	}
}

func checkForOtherCommands(args []string) {
	if len(args) == 0 {
		return
	}
	cmd, ok := commands[args[0]]
	if !ok {
		return
	}
	cmd()
}
