package usecases

import (
	"github.com/miodzie/dong"
)

type RandomDong struct {
	repository dong.Repository
}

func NewRandomDong(repo dong.Repository) *RandomDong {
	return &RandomDong{repository: repo}
}

type RandomDongRequest struct {
	Category string
}

type RandomDongResponse struct {
	Emoji string
}

func (r RandomDong) Pick(req RandomDongRequest) (RandomDongResponse, error) {
	if req.Category != "" {
		emoji, err := r.repository.RandomByCategory(req.Category)
		return RandomDongResponse{Emoji: emoji.Text}, err
	}

	emoji, err := r.repository.Random()
	return RandomDongResponse{Emoji: emoji.Text}, err
}
