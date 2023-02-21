package executors

import (
	"context"
	"time"

	"github.com/devNica/mochileros/commons"
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
	repositories.FileRepo
}

func NewUserSrvExecutor(userRepo *repositories.UserRepo, fileRepo *repositories.FileRepo) services.UserService {
	return &userServiceExecutor{UserRepo: *userRepo, FileRepo: *fileRepo}
}

func (srv *userServiceExecutor) RegisterKYC(
	ctx context.Context,
	kyc request.KYCRequestModel,
	newFiles []request.FileRequestModel) {

	userId, err := uuid.Parse(kyc.UserId)
	exceptions.PanicLogging(err)

	kycEntity := entities.UserInfo{
		FirstName: kyc.FirstName,
		LastName:  kyc.LastName,
		Address:   kyc.Address,
		Birthdate: kyc.Birthdate,
		DNI:       kyc.DNI,
		CountryID: kyc.CountryId,
		UserId:    userId,
	}

	commons.ValidateModel(kycEntity)

	registerError := srv.UserRepo.InsertKYC(ctx, kycEntity)

	if registerError != nil {
		exceptions.PanicLogging(registerError)
	}

	for _, newFile := range newFiles {

		file := entities.File{
			Filename:  uuid.New(),
			Filetype:  newFile.Filetype,
			Filesize:  newFile.Filesize,
			Binary:    newFile.Buffer,
			CreatedAt: time.Now(),
		}

		srv.FileRepo.InsertAssetByUserId(
			ctx,
			file,
			newFile.AssetTypeId,
			userId,
		)
	}

}

func (srv *userServiceExecutor) ChangeAccountStatus(ctx context.Context, userId string) response.UserResponseModel {
	response, err := srv.UserRepo.UpdateUserAccountStatus(ctx, userId)
	exceptions.PanicLogging(err)
	return response
}
