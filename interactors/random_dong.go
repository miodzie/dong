package interactors

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
	Category string // Optional
}

type RandomDongResp struct {
	Emoji string
	Error error
}

func (r RandomDong) Handle(req RandomDongReq) RandomDongResp {
	var ding dong.Emoji
	var err error
	// Maybe just bake this in with the random, I don't like the if else.
	if req.Category != "" {
		ding, err = r.repository.RandomByCategory(req.Category)
	} else {
		ding, err = r.repository.Random()
	}

	return RandomDongResp{
		Emoji: ding.Text,
		Error: err,
	}
}
