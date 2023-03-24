package controller

import (
	"net/http"
	"time"

	"github.com/example/gin_framework/dto"
	"github.com/example/gin_framework/helper"
	"github.com/example/gin_framework/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(c *gin.Context)
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService,
		jwtService,
	}
}

func (c *loginController) Login(ctx *gin.Context) {
	var credentials dto.Credentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse(err))
		return
	}

	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)

	if isAuthenticated {
		token, err := c.jwtService.GenerateJWT(credentials.Username, 15*time.Minute)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, token)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.ErrorResponse(helper.ErrInvalidToken))
}
