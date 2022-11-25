package phrase

import (
	"context"
	"server/internal/dao"
	"server/internal/model"
	"server/internal/model/do"
	"server/internal/service"
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
func (s *sPhrase) Query(ctx context.Context, in model.PhraseQueryInput) (PhraseExplain string, err error) {
	phrase := in.QueryString
	explain := QueryPhrase(phrase)

	_, err = dao.Phrase.Ctx(ctx).Data(do.Phrase{
		Phrase:  phrase,
		Explain: explain,
	}).Insert()

	return explain, nil
}
