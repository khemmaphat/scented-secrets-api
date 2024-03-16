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

func ApplyUserHandler(r *gin.Engine, client *firestore.Client) {

	userRepo := repository.MakeUserRepository(client)
	userService := service.MakeUserService(userRepo)
	userHandler := MakeUserHandler(userService)

	userGroup := r.Group("/api")
	{
		userGroup.GET("/user", userHandler.GetUserById)
		userGroup.POST("/user", userHandler.CreateUser)
		userGroup.POST("/login", userHandler.LoginUser)
		userGroup.PATCH("/edituser", userHandler.EditUser)
	}

}

type UserHandler struct {
	userService infService.IUserService
}

func MakeUserHandler(userService infService.IUserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h UserHandler) GetUserById(c *gin.Context) {
	id := c.Query("id")
	var res model.HTTPResponse
	var user entities.User

	if id == "" {
		res.SetError("Require user id", 400, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	user, err := h.userService.GetUserById(context.Background(), id)
	if err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetSuccess("Get user success", 200, user)
	c.JSON(http.StatusOK, res)
}

func (h UserHandler) CreateUser(c *gin.Context) {
	var user entities.User
	var res model.HTTPResponse

	if err := c.ShouldBind(&user); err != nil {
		res.SetError(err.Error(), 400, err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := h.userService.CrateUser(context.Background(), user); err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetSuccess("Crete user success", 200, user)
	c.JSON(http.StatusOK, res)
}

func (h UserHandler) LoginUser(c *gin.Context) {
	var user entities.User
	var res model.HTTPResponse

	if err := c.Bind(&user); err != nil {
		res.SetError(err.Error(), 400, err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	id, err := h.userService.LoginUser(context.Background(), user)
	if err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}
	res.SetSuccess(id, 200, user)
	c.JSON(http.StatusOK, res)
}

func (h UserHandler) EditUser(c *gin.Context) {
	id := c.Query("id")
	var user entities.User
	var res model.HTTPResponse

	if id == "" {
		res.SetError("Require user id", 400, nil)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.Bind(&user); err != nil {
		res.SetError(err.Error(), 400, err)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	err := h.userService.EditUser(context.Background(), id, user)
	if err != nil {
		res.SetError(err.Error(), 200, err)
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetSuccess("Edit user success", 200, user)
	c.JSON(http.StatusOK, res)
}
