package service

import (
	"context"
	"strings"

	"github.com/khemmaphat/scented-secrets-api/src/entities"
	infRepo "github.com/khemmaphat/scented-secrets-api/src/repository/infRepo"
	infServ "github.com/khemmaphat/scented-secrets-api/src/service/infService"
)

type QuestionService struct {
	questionRepository infRepo.IQuestionRepository
	perfumeRepository  infRepo.IPerfumeRepository
}

func MakeQuestionService(questionRepository infRepo.IQuestionRepository, perfumeRepository infRepo.IPerfumeRepository) infServ.IQuestionService {
	return &QuestionService{questionRepository: questionRepository, perfumeRepository: perfumeRepository}
}

func (s QuestionService) GetQuestions(ctx context.Context) ([]entities.Question, error) {
	questions, err := s.questionRepository.GetQuestions(ctx)
	return questions, err
}

func (s QuestionService) GetResultQuestion(ctx context.Context, answered string) (entities.ResultQuestion, error) {
	cosineValue := s.questionRepository.GetResultQuestion(answered)
	perfumeId, perfumeDetail, err := s.perfumeRepository.GetPerfumeByCosineValue(ctx, cosineValue)

	if err != nil {
		return entities.ResultQuestion{}, err
	}

	PerfumeNotesTemp := strings.Split(perfumeDetail.Notes, ",")
	var PerfumeNotes entities.Notes
	for i, pt := range PerfumeNotesTemp {
		if i < 2 {
			PerfumeNotes.BaseNotes = append(PerfumeNotes.BaseNotes, pt)
		} else if i >= 2 && i <= 4 {
			PerfumeNotes.TopNotes = append(PerfumeNotes.TopNotes, pt)
		} else {
			PerfumeNotes.MiddleNotes = append(PerfumeNotes.MiddleNotes, pt)
		}

	}

	resultQuetion := entities.ResultQuestion{
		PerfumeId:   perfumeId,
		Name:        perfumeDetail.Name,
		Brand:       perfumeDetail.Brand,
		Notes:       PerfumeNotes,
		Description: perfumeDetail.Description,
		ImfUrl:      perfumeDetail.ImgUrl,
	}

	return resultQuetion, nil
}
