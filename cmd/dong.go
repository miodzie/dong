package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"os/user"
	"path"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/miodzie/dong/core"
	"github.com/spf13/cobra"
)

var db *gorm.DB

func init() {
	usr, err := user.Current()
	workDir := path.Join(usr.HomeDir, ".dong")

	if _, err := os.Stat(workDir); os.IsNotExist(err) {
		err := os.Mkdir(workDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	db, err = gorm.Open("sqlite3", path.Join(workDir, "dongs.db"))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	err = core.Initialize(db)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&core.Dong{})
}

var dongCmd = &cobra.Command{
	Args:  cobra.ArbitraryArgs,
	Use:   "dong",
	Short: "Print a random dong.",
	Long:  `Print a random dong.`,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		qb := db.Model(&core.Dong{}).Select("id")
		if len(args) != 0 {
			qb = qb.Where("category IN (?)", args)
		}
		rows, err := qb.Rows()
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var id int
			rows.Scan(&id)
			ids = append(ids, id)
		}

		if len(ids) == 0 {
			fmt.Println("No dongers. Scrape some dongs with: 'dong scrape'")
			return
		}

		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(len(ids))
		var id int
		for i, num := range ids {
			if i == random {
				id = num
				break
			}
		}
		qb = db.Model(&core.Dong{}).Where("id = ?", id)

		if len(args) != 0 {
			qb = qb.Where("category IN (?)", args)
		}

		var dong core.Dong
		qb.First(&dong)
		fmt.Println(dong)
	},
}

func Execute() {
	if err := dongCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
