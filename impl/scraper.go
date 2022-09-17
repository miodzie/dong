package impl

import (
	"fmt"
	"github.com/miodzie/dong"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const DONGERLIST = "http://dongerlist.com"

func NewScraper() *WebFetcher {
	return &WebFetcher{domain: DONGERLIST}
}

type WebFetcher struct {
	domain string
}

func (s *WebFetcher) Fetch() ([]dong.Emoji, error) {
	var dongs []dong.Emoji
	doc, err := s.fetchDocument(s.domain)
	if err != nil {
		return dongs, err
	}
	categories := findCategories(doc)

	for _, category := range categories {
		fmt.Println("Scraping: " + s.domain + "/category/" + category)
		page, err := s.fetchDocument(s.domain + "/category/" + category)
		if err != nil {
			return dongs, err
		}
		tot := page.Find(".last").First().Text()
		if tot == "" {
			tot = "1"
		}

		totalPages, err := strconv.Atoi(tot)
		fmt.Println("TOTS:" + strconv.Itoa(totalPages))
		if err != nil {
			return dongs, err
		}

		for i := 1; i <= totalPages; i++ {
			if i == 1 {
				page.Find(".donger").Each(func(i int, dng *goquery.Selection) {
					if dng.Text() != "" {
						emoji := dong.Emoji{}
						emoji.Text = dng.Text()
						emoji.Category = category
						dongs = append(dongs, emoji)
						fmt.Printf("new dong: %s !!\n", emoji)
					}
				})
			}
		}

	}

	return dongs, nil
}

func findCategories(doc *goquery.Document) []string {
	var categories []string
	doc.Find(".list-2-anchor").Each(func(i int, selection *goquery.Selection) {
		category := selection.AttrOr("href", "")
		if category != "" {
			split := strings.Split(category, "/")
			categories = append(categories, split[len(split)-1])
		}
	})

	return categories
}

func (s *WebFetcher) fetchDocument(url string) (*goquery.Document, error) {
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
