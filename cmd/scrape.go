package cmd

import (
	"fmt"

	"github.com/miodzie/dong/core"
	"github.com/spf13/cobra"
)

func init() {
	dongCmd.AddCommand(scrapeCmd)
}

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape dongers.",
	Long:  `ヽ༼ຈل͜ຈ༽ﾉ FOREVER DONG ヽ༼ຈل͜ຈ༽ﾉ`,
	Run: func(cmd *cobra.Command, args []string) {
		scraper := &core.Scraper{Domain: "http://dongerlist.com"}
		err := scraper.Run()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	},
}
