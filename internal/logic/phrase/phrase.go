package phrase

import (
	"context"
	"server/internal/model"
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
	phrase := QueryPhrase(in.QueryString)
	return phrase, nil
}
