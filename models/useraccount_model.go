package models

import (
	"time"
)

type UserAccounRequestModel struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponseModel struct {
	Email     string    `json:"email"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

type KYCRequestModel struct {
	UserId    string `json:"userId" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
