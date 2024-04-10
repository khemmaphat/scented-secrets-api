package handler

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/khemmaphat/scented-secrets-api/src/entities"
	"github.com/khemmaphat/scented-secrets-api/src/model"
	"github.com/khemmaphat/scented-secrets-api/src/repository"
	"github.com/khemmaphat/scented-secrets-api/src/service"
	"github.com/khemmaphat/scented-secrets-api/src/service/infService"
)

func ApplyPerfumeHandler(r *gin.Engine, client *firestore.Client) {

	perfumeRepo := repository.MakePerfumeRepository(client)
	perfumeService := service.MakePerfumeService(perfumeRepo)
	perfumeHandler := MakePerfumeHandler(perfumeService)

	perfumeGroup := r.Group("/api")
	{
		perfumeGroup.GET("/perfume", perfumeHandler.GetPerfumeById)
		perfumeGroup.POST("/createperfume", perfumeHandler.AddPerfumeData)
		perfumeGroup.POST("/searchperfume", perfumeHandler.SearchPerfumePagination)
		perfumeGroup.POST("/createnote", perfumeHandler.AddNotesData)
		perfumeGroup.GET("/getallgroupnote", perfumeHandler.GetAllNoteGroup)

	}

}

type PerfumeHandler struct {
	perfumeService infService.IPerfumeService
}

func MakePerfumeHandler(perfumeService infService.IPerfumeService) *PerfumeHandler {
	return &PerfumeHandler{perfumeService: perfumeService}
}

func (h PerfumeHandler) GetPerfumeById(c *gin.Context) {
	id := c.Query("id")
	var perfume model.PerfumeDetail
	var res model.HTTPResponse

	if id == "" {
		res.SetError("Require user id", 400, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	perfume, err := h.perfumeService.GetPerfumeById(context.Background(), id)
	if err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetSuccess("get perfume success", 200, perfume)
	c.JSON(http.StatusOK, res)
}

func (h PerfumeHandler) AddPerfumeData(c *gin.Context) {
	var perfumes []entities.Perfume
	var res model.HTTPResponse

	if err := c.ShouldBind(&perfumes); err != nil {
		res.SetError(err.Error(), 400, err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := h.perfumeService.AddPerfumeData(perfumes); err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetSuccess("Crete perfume success", 200, perfumes)
	c.JSON(http.StatusOK, res)
}

func (h PerfumeHandler) SearchPerfumePagination(c *gin.Context) {
	var req entities.PerfumePaginationRequest
	var res model.HTTPResponse

	if err := c.Bind(&req); err != nil {
		res.SetError(err.Error(), 400, err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	data, total, err := h.perfumeService.SearchPerfumePagination(context.Background(), req)
	if err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.Total = total
	res.SetSuccess("Search Success", 200, data)
	c.JSON(http.StatusOK, res)
}

func (h PerfumeHandler) AddNotesData(c *gin.Context) {
	var notes []entities.Note
	var res model.HTTPResponse

	if err := c.ShouldBind(&notes); err != nil {
		res.SetError(err.Error(), 400, err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := h.perfumeService.AddNoteData(notes); err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetSuccess("Crete Notes success", 200, notes)
	c.JSON(http.StatusOK, res)
}

func (h PerfumeHandler) GetAllNoteGroup(c *gin.Context) {
	var res model.HTTPResponse
	groupNote, err := h.perfumeService.GetAllNoteGroup(context.Background())

	if err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetSuccess("Get Group Note success", 200, groupNote)
	c.JSON(http.StatusOK, res)
}