package response

import (
	"time"

	"github.com/google/uuid"
)

type UserResponseModel struct {
	Id        uuid.UUID
	Email     string
	IsActive  bool
	CreatedAt time.Time
}

type UserInfoResponseModel struct {
	Id       string
	Email    string
	IsActive bool
	UserInfo struct {
		FirstName string
		LastName  string
	}
	Profile   []string
	CreatedAt time.Time
}

type LoginResponseModel struct {
	Id       string
	Email    string
	IsActive bool
	UserInfo struct {
		FirstName string
		LastName  string
	}
	Token     string
	CreatedAt time.Time
}
