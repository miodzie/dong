package interactors

import "github.com/miodzie/dong/domain"

type RandomDong struct {
	random domain.RandomDonger
}

type RandomDongReq struct {
	Category string // Optional
}

type RandomDongResp struct {
	Emoji string
	Error error
}

func (r RandomDong) Handle(req RandomDongReq) RandomDongResp {
	var dong domain.Dong
	var err error
	// Maybe just bake this in with the random, I don't like the if else.
	if req.Category != "" {
		dong, err = r.random.DongByCategory(req.Category)
	} else {
		dong, err = r.random.Dong()
	}

	return RandomDongResp{
		Emoji: dong.Emoji,
		Error: err,
	}
}
