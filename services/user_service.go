package services

import (
	"context"

	"github.com/devNica/mochileros/dto/request"
)

type UserService interface {
	RegisterKYC(ctx context.Context, kyc request.KYCRequestModel, newFiles []request.FileRequestModel)
	ChangeAccountStatus(ctx context.Context, userId string, statusId uint8)
}
