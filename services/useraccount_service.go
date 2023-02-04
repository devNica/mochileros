package services

import (
	"context"

	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
)

type UserAccountService interface {
	UserAccountRegister(ctx context.Context, newUser request.UserAccounRequestModel)
	RegisterKYC(ctx context.Context, kyc request.KYCRequestModel)
	ChangeAccountStatus(ctx context.Context, userId string) response.UserResponseModel
	GetUserByEmail(ctx context.Context, email string) response.UserResponseModel
	GetCompleteUserInfo(ctx context.Context, userId string) response.UserInfoResponseModel
}
