package service

import (
	"fmt"

	"github.com/example/gin_framework/helper"
	"github.com/example/gin_framework/initializers"
	"github.com/example/gin_framework/model"
)

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct{}

func NewLoginService() LoginService {
	return &loginService{}
}

func (s *loginService) Login(username, password string) bool {
	user := model.User{}
	initializers.DB.Where(model.User{Username: username}).Find(&user)
	fmt.Println(username, password)
	err := helper.CheckPassword(password, user.Password)
	return err == nil

}
