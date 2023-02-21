package services

import (
	"context"

	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
)

type AuthService interface {
	CustomerRegister(ctx context.Context, newCustomer request.UserAccounRequestModel)
	UserLogin(ctx context.Context, user request.UserAccounRequestModel) (login response.LoginResponseModel)
	RegisterKYC(ctx context.Context, kyc request.KYCRequestModel)
	ChangeAccountStatus(ctx context.Context, userId string) response.UserResponseModel
}
