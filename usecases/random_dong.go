package usecases

import (
	"github.com/miodzie/dong"
)

type RandomDong struct {
	repository dong.Repository
}

func NewRandomDongInteractor(repo dong.Repository) *RandomDong {
	return &RandomDong{repository: repo}
}

type RandomDongReq struct {
	Category string
}

type RandomDongResp struct {
	Emoji string
	Error error
}

func (r RandomDong) Handle(req RandomDongReq) RandomDongResp {
	var emoji dong.Emoji
	var resp RandomDongResp
	if req.Category != "" {
		emoji, resp.Error = r.repository.RandomByCategory(req.Category)
		return RandomDongResp{Emoji: emoji.Text}
	}

	emoji, resp.Error = r.repository.Random()
	return RandomDongResp{Emoji: emoji.Text}
}
