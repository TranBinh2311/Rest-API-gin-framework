package dto

import (
	"time"

	"github.com/example/gin_framework/model"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum" validate:"is-contain"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func NewUserResponse(user model.User) UserResponse {
	return UserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

type GetAccountParams struct {
	ID int64 `uri:"id" binding:"required"`
}

type ListAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required"`
	PageSize int32 `form:"page_size" binding:"required"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
