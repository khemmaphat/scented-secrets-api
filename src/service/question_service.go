package service

import (
	"context"

	"github.com/khemmaphat/scented-secrets-api/src/entities"
	infRepo "github.com/khemmaphat/scented-secrets-api/src/repository/infRepo"
	infServ "github.com/khemmaphat/scented-secrets-api/src/service/infService"
)

type QuestionService struct {
	questionRepository infRepo.IQuestionRepository
}

func MakeQuestionService(questionRepository infRepo.IQuestionRepository) infServ.IQuestionService {
	return &QuestionService{questionRepository: questionRepository}
}

func (s QuestionService) GetQuestions(ctx context.Context) ([]entities.Question, error) {
	questions, err := s.questionRepository.GetQuestions(ctx)
	return questions, err
}
