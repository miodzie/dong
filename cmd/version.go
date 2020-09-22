package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	dongCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of dong",
	Long:  `ヽ༼ຈل͜ຈ༽ﾉ FOREVER DONG ヽ༼ຈل͜ຈ༽ﾉ`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ヽ༼ຈل͜ຈ༽ﾉ FOREVER DONG ヽ༼ຈل͜ຈ༽ﾉ")
	},
}
