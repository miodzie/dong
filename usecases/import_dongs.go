package usecases

import (
	"fmt"
	"github.com/miodzie/dong"
)

type ImportDongs struct {
	fetcher    dong.Fetcher
	repository dong.Repository
}

func NewImportDongs(
	fetcher dong.Fetcher,
	repository dong.Repository) *ImportDongs {
	return &ImportDongs{
		fetcher:    fetcher,
		repository: repository,
	}
}

type ImportDongsResponse struct {
	Message string
	Error   error
}

func (s ImportDongs) Import() ImportDongsResponse {
	dongs, err := s.fetcher.Fetch()
	if err != nil {
		return ImportDongsResponse{Error: err}
	}

	err = s.repository.Save(dongs)

	return ImportDongsResponse{
		Message: fmt.Sprintf("%d new dongs created!", len(dongs)),
		Error:   nil,
	}
}
