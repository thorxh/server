// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"gitee.com/thorxh/thor-api/internal/model"
)

type (
	IPhrase interface {
		Query(ctx context.Context, in model.PhraseQueryInput) (PhraseExplain string, err error)
	}
)

var (
	localPhrase IPhrase
)

func Phrase() IPhrase {
	if localPhrase == nil {
		panic("implement not found for interface IPhrase, forgot register?")
	}
	return localPhrase
}

func RegisterPhrase(i IPhrase) {
	localPhrase = i
}
