package models

import (
	"time"

	"github.com/google/uuid"
)

type UserAccounRequestModel struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type KYCRequestModel struct {
	UserId    string `json:"userId" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type CompleteUserRequestModel struct {
	Id        uuid.UUID
	Email     string
	FirstName string
	LastName  string
	IsActive  bool
	UserId    uuid.UUID
}

type UserResponseModel struct {
	Email     string    `json:"email"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

type KYCResponseModel struct {
	Id       uuid.UUID
	Email    string
	IsActive bool
	KYC      struct {
		FirstName string
		LastName  string
	}
}

type UpdateUserAccountStatusResModel struct {
	UserId   uuid.UUID
	IsActive bool
}
