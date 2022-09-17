package interactors

import (
	"fmt"
	"github.com/miodzie/dong"
)

type Scrape struct {
	fetcher    dong.Fetcher
	repository dong.Repository
}

func NewScrapeDongsInteractor(
	fetcher dong.Fetcher,
	repository dong.Repository) *Scrape {
	return &Scrape{
		fetcher:    fetcher,
		repository: repository,
	}
}

type ScrapeResp struct {
	Message string
	Error   error
}

func (s Scrape) Handle() ScrapeResp {
	dongs, err := s.fetcher.Fetch()
	if err != nil {
		return ScrapeResp{Error: err}
	}

	err = s.repository.Save(dongs)

	return ScrapeResp{
		Message: fmt.Sprintf("%d new dongs created!", len(dongs)),
		Error:   nil,
	}
}
