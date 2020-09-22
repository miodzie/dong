package cmd

import (
	"fmt"
	"strings"

	"github.com/miodzie/dong/core"
	"github.com/spf13/cobra"
)

func init() {
	dongCmd.AddCommand(catCmd)
}

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Shows available categories.",
	Long:  `Shows available categories.`,
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := db.Model(&core.Dong{}).Select("category").Group("category").Rows()
		if err != nil {
			panic(err)
		}
		var cat []string
		for rows.Next() {
			var cate string
			rows.Scan(&cate)
			cat = append(cat, cate)
		}
		fmt.Println(strings.Join(cat[1:], ", "))
	},
}
