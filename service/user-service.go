package service

import (
	"github.com/example/gin_framework/dto"
	"github.com/example/gin_framework/helper"
	"github.com/example/gin_framework/initializers"
	"github.com/example/gin_framework/model"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest) (dto.UserResponse, error)
	GetAllUsers(req dto.ListAccountsRequest) ([]dto.UserResponse, error)
}

type userService struct {
}

func New() UserService {
	return &userService{}
}

func (service *userService) CreateUser(req dto.CreateUserRequest) (dto.UserResponse, error) {
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}

	dataEntry := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
	}

	result := initializers.DB.Create(&dataEntry)
	err = result.Error

	covertedData := dto.NewUserResponse(dataEntry)
	return covertedData, err
}

func (service *userService) GetAllUsers(req dto.ListAccountsRequest) ([]dto.UserResponse, error) {
	listUser := make([]model.User, 0)
	covertedData := make([]dto.UserResponse, 0)

	pageID, pageSize := int(req.PageID), int(req.PageSize)
	result := initializers.DB.Table("users").Offset((pageID - 1) * pageSize).Limit(pageSize).Find(&listUser)

	err := result.Error

	for _, user := range listUser {
		covertedData = append(covertedData, dto.NewUserResponse(user))
	}
	return covertedData, err
}
