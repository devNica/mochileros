package services

import (
	"context"

	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
)

type UserService interface {
	RegisterKYC(ctx context.Context, kyc request.KYCRequestModel, newFiles []request.FileRequestModel)
	ChangeAccountStatus(ctx context.Context, userId string) response.UserResponseModel
}
