package phrase

import (
	"context"
	"server/internal/model"
	"server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sPhrase struct{}
)

func init() {
	service.RegisterPhrase(New())
}

func New() *sPhrase {
	return &sPhrase{}
}

// Query query phrase.
func (s *sPhrase) Query(ctx context.Context, in model.PhraseQueryInput) (err error) {
	phrase := QueryPhrase(in.QueryString)
	return gerror.Newf(`%s: %s`, in.QueryString, phrase)
}
