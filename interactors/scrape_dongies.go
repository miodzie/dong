package interactors

import "github.com/miodzie/dong"

type Scrape struct {
	fetcher    dong.Fetcher
	repository dong.Repository
}

type ScrapeResp struct {
	Error error
}

func (s Scrape) Handle() ScrapeResp {
	dongs, err := s.fetcher.Fetch()
	if err != nil {
		return ScrapeResp{Error: err}
	}

	err = s.repository.Save(dongs)

	return ScrapeResp{
		Error: nil,
	}
}
