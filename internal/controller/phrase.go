package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"server/api/v1"
)

var (
	Phrase = cPhrase{}
)

type cPhrase struct{}

func (c *cPhrase) Phrase(ctx context.Context, req *v1.PhraseReq) (res *v1.PhraseRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
