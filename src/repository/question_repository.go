package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/khemmaphat/scented-secrets-api/src/entities"
	infRepo "github.com/khemmaphat/scented-secrets-api/src/repository/infRepo"
	"google.golang.org/api/iterator"
)

type QuestionRepository struct {
	client *firestore.Client
}

func MakeQuestionRepository(client *firestore.Client) infRepo.IQuestionRepository {
	return &QuestionRepository{client: client}
}

func (r QuestionRepository) GetQuestions(ctx context.Context) ([]entities.Question, error) {
	var questions []entities.Question

	iter := r.client.Collection("questions").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var question entities.Question
		if err := doc.DataTo(&question); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}
