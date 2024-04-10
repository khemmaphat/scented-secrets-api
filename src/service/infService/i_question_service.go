package infService

import (
	"context"

	"github.com/khemmaphat/scented-secrets-api/src/entities"
)

type IQuestionService interface {
	GetQuestions(ctx context.Context) ([]entities.Question, error)
}