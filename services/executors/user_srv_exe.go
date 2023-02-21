package executors

import (
	"context"

	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
	"github.com/google/uuid"
)

type userServiceExecutor struct {
	repositories.UserRepo
}

func NewUserSrvExecutor(repo *repositories.UserRepo) services.UserService {
	return &userServiceExecutor{UserRepo: *repo}
}

func (srv *userServiceExecutor) RegisterKYC(ctx context.Context, kyc request.KYCRequestModel) {

	userId, err := uuid.Parse(kyc.UserId)
	exceptions.PanicLogging(err)

	kycEntity := entities.UserInfo{
		FirstName: kyc.FirstName,
		LastName:  kyc.LastName,
		UserId:    userId,
	}

	srv.UserRepo.InsertKYC(ctx, kycEntity)
}

func (srv *userServiceExecutor) ChangeAccountStatus(ctx context.Context, userId string) response.UserResponseModel {
	response, err := srv.UserRepo.UpdateUserAccountStatus(ctx, userId)
	exceptions.PanicLogging(err)
	return response
}
