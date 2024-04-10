package infRepo

import (
	"context"

	"github.com/khemmaphat/scented-secrets-api/src/entities"
)

type IQuestionRepository interface {
	GetQuestions(ctx context.Context) ([]entities.Question, error)
}
