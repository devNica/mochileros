package services

import (
	"context"

	"github.com/devNica/mochileros/models"
)

type UserAccountService interface {
	UserAccountRegister(ctx context.Context, model models.UserAccounRequestModel)
	RegisterKYC(ctx context.Context, model models.KYCRequestModel)
	GetUserByEmail(ctx context.Context, email string) models.UserResponseModel
}
