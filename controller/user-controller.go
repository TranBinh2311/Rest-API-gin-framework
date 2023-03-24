package controller

import (
	"fmt"
	"net/http"

	"github.com/example/gin_framework/dto"
	"github.com/example/gin_framework/helper"
	"github.com/example/gin_framework/service"
	"github.com/example/gin_framework/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserController interface {
	GetAllUsers(c *gin.Context)
	Create(ctx *gin.Context)
	HeathCheck(ctx *gin.Context)
}

type controller struct {
	service service.UserService
}

var validate *validator.Validate

func New(service service.UserService) UserController {
	validate = validator.New()
	validate.RegisterValidation("is-contain", validators.ValidateContains)
	return &controller{
		service,
	}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} HeathCheck
// @Router /user/heathCheck [get]
func (c *controller) HeathCheck(ctx *gin.Context) {
	//
	ctx.JSON(http.StatusOK, "This is heath check api from user controller !")
}

func (c *controller) Create(ctx *gin.Context) {
	//
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err))
		return
	}
	err := validate.Struct(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err))
		return
	}
	//
	res, err := c.service.CreateUser(req)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse(err))
		return
	}

	ctx.JSON(200, res)
}

func (c *controller) GetAllUsers(ctx *gin.Context) {
	var queries dto.ListAccountsRequest
	if err := ctx.ShouldBindQuery(&queries); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.service.GetAllUsers(queries)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse(err))
		return
	}
	ctx.JSON(200, res)
}
