package core

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	Domain     string
	Categories []string
	Limit      uint
}

func (s *Scraper) Run() error {
	doc, err := FetchDocument(s.Domain)
	if err != nil {
		return err
	}
	if len(s.Categories) == 0 {
		doc.Find(".list-2-anchor").Each(func(i int, selection *goquery.Selection) {
			category := selection.AttrOr("href", "")
			if category != "" {
				split := strings.Split(category, "/")
				s.Categories = append(s.Categories, split[len(split)-1])
			}
		})
	}

	for _, cat := range s.Categories {
		fmt.Println("Scraping: " + s.Domain + "/category/" + cat)
		page, err := FetchDocument(s.Domain + "/category/" + cat)
		if err != nil {
			return err
		}
		tot := page.Find(".last").First().Text()
		if tot == "" {
			tot = "1"
		}

		totalPages, err := strconv.Atoi(tot)
		fmt.Println("TOTS:" + strconv.Itoa(totalPages))
		if err != nil {
			return err
		}

		for i := 1; i <= totalPages; i++ {
			if i == 1 {
				page.Find(".donger").Each(func(i int, dng *goquery.Selection) {
					if dng.Text() != "" {
						dong := &Dong{}
						db.Where("dong = ?", dong.Dong).First(&dong)
						if dong.Dong == "" {
							dong.Dong = dng.Text()
							dong.Category = cat
							db.Create(&dong)
							fmt.Println("New dong created: " + dong.Dong)
						}
					}
				})
			}
		}

	}

	return nil
}

// Some helper method ...
func FetchDocument(url string) (*goquery.Document, error) {
	r, err := http.Get(url)

	if err != nil {
		return nil, nil
	}

	doc, docerr := goquery.NewDocumentFromReader(r.Body)

	if docerr != nil {
		return nil, nil
	}
	return doc, nil
}
