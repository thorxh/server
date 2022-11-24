package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PhraseReq struct {
	g.Meta `path:"/phrase" tags:"Phrase" method:"get" summary:"query phrase"`
}
type PhraseRes struct {
	g.Meta `mime:"text/html" example:"string"`
}