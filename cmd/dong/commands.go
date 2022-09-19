package main

import (
	"errors"
	"fmt"
	"github.com/miodzie/dong/impl"
	"github.com/miodzie/dong/usecases"
	"strings"
)

type command struct {
	cmd         string
	description string
	handle      func(args []string) error
}

var commands map[string]command

func init() {
	commands = make(map[string]command)
	commands["count"] = command{
		cmd:         "count",
		description: "Display the amount of dongs available.",
		handle: func(args []string) error {
			fmt.Println(repository.Count())
			return nil
		},
	}
	commands["cat"] = command{
		cmd:         "cat",
		description: "Display available categories to use. e.g. `dong happy`",
		handle: func(args []string) error {
			cats, err := repository.Categories()
			if err != nil {
				return err
			}
			fmt.Println(strings.Join(cats, ", "))

			return nil
		},
	}
	commands["scrape"] = command{
		cmd:         "scrape",
		description: "Web scrape fresh dongs off the press from dongerlist.com",
		handle: func(args []string) error {
			exec := usecases.NewScrapeDongsInteractor(impl.NewScraper(), repository)
			resp := exec.Handle()
			if resp.Error != nil {
				return resp.Error
			}
			fmt.Println(resp.Message)

			return nil
		},
	}
	commands["version"] = command{
		cmd:         "version",
		description: "Display the version of the dong program.",
		handle: func(args []string) error {
			fmt.Println("ヽ༼ຈل͜ຈ༽ﾉ FOREVER DONG ヽ༼ຈل͜ຈ༽ﾉ")
			return nil
		},
	}
	commands["help"] = command{
		cmd:         "help",
		description: "Display helpful resources?",
		handle: func(args []string) error {
			if len(args) > 1 {
				if cmd, ok := commands[args[1]]; ok {
					fmt.Println(cmd.description)
					return nil
				}
				return errors.New(fmt.Sprintf("Unknown command: `%s`\n", args[1]))
			}
			var keys []string
			for k := range commands {
				keys = append(keys, k)
			}
			fmt.Println(strings.Join(keys, ", "))
			fmt.Println("Type `help command` for more info.")
			return nil
		},
	}
}

func handleCommands(args []string) error {
	if len(args) == 0 {
		return printRandomDong(args)
	}
	cmd, ok := commands[args[0]]
	if !ok {
		return printRandomDong(args)
	}
	return cmd.handle(args)
}

func printRandomDong(args []string) error {
	req := usecases.RandomDongReq{}
	if len(args) > 0 {
		req.Category = args[0]
	}
	controller := usecases.NewRandomDongInteractor(repository)
	resp := controller.Handle(req)
	if resp.Error != nil {
		return resp.Error
	}

	fmt.Println(resp.Emoji)

	return nil
}
