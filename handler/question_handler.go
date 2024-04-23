package handler

import (
	"context"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/khemmaphat/scented-secrets-api/src/model"
	"github.com/khemmaphat/scented-secrets-api/src/repository"
	"github.com/khemmaphat/scented-secrets-api/src/service"
	"github.com/khemmaphat/scented-secrets-api/src/service/infService"
)

func ApplyQuestionHandler(r *gin.Engine, client *firestore.Client) {

	questionRepo := repository.MakeQuestionRepository(client)
	perfumeRepo := repository.MakePerfumeRepository(client)
	questionService := service.MakeQuestionService(questionRepo, perfumeRepo)
	questionHandler := MakeQuestionHandler(questionService)

	questionGroup := r.Group("/api")
	{
		questionGroup.GET("/question", questionHandler.GetQuestions)
		questionGroup.GET("/resultquestion", questionHandler.GetResultQuestion)
	}

}

type QuestionHandler struct {
	questionService infService.IQuestionService
}

func MakeQuestionHandler(questionService infService.IQuestionService) *QuestionHandler {
	return &QuestionHandler{questionService: questionService}
}

func (h QuestionHandler) GetQuestions(c *gin.Context) {
	var res model.HTTPResponse

	question, err := h.questionService.GetQuestions(context.Background())
	if err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetSuccess("get question success", 200, question)
	c.JSON(http.StatusOK, res)
}

func (h QuestionHandler) GetResultQuestion(c *gin.Context) {
	var res model.HTTPResponse
	answered := c.Query("answered")
	answered = strings.ReplaceAll(answered, "%20", " ")

	resultQuestion, err := h.questionService.GetResultQuestion(context.Background(), answered)
	if err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetSuccess("get result question success", 200, resultQuestion)
	c.JSON(http.StatusOK, res)
}
