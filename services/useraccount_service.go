package services

import (
	"context"

	"github.com/devNica/mochileros/models"
)

type UserAccountService interface {
	UserAccountRegister(ctx context.Context, model models.UserAccounRequestModel)
	GetUserByEmail(ctx context.Context, email string) models.UserResponseModel
}
