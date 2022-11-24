package controller

import (
	"context"
	"server/internal/model"
	"server/internal/service"

	"server/api/v1"
)

var (
	Phrase = cPhrase{}
)

type cPhrase struct{}

func (c *cPhrase) Phrase(ctx context.Context, req *v1.PhraseReq) (res *v1.PhraseRes, err error) {
	var PhraseExplain string
	PhraseExplain, err = service.Phrase().Query(ctx, model.PhraseQueryInput{
		QueryString: req.QueryString,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.PhraseRes{
		Phrase:  req.QueryString,
		Explain: PhraseExplain,
	}
	return
}
