package response

import (
	"time"

	"github.com/google/uuid"
)

type UserResponseModel struct {
	Id            uuid.UUID
	Email         string
	TwoFactorAuth bool
	Status        string
	CreatedAt     time.Time
}

type UserInfoResponseModel struct {
	Id            string
	Email         string
	Password      string
	TwoFactorAuth bool
	UserInfo      struct {
		FirstName string
		LastName  string
	}
	Profile   []string
	CreatedAt time.Time
}

type LoginResponseModel struct {
	Id            string
	Email         string
	TwoFactorAuth bool
	UserInfo      struct {
		FirstName string
		LastName  string
	}
	Token     string
	CreatedAt time.Time
}
