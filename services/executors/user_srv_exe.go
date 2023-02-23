package executors

import (
	"context"
	"time"

	"github.com/devNica/mochileros/commons"
	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
	"github.com/google/uuid"
)

type userServiceExecutor struct {
	repositories.UserRepo
	repositories.FileRepo
	repositories.BackofficeRepo
}

func NewUserSrvExecutor(
	userRepo *repositories.UserRepo,
	fileRepo *repositories.FileRepo,
	backofficeRepo *repositories.BackofficeRepo) services.UserService {
	return &userServiceExecutor{
		UserRepo:       *userRepo,
		FileRepo:       *fileRepo,
		BackofficeRepo: *backofficeRepo,
	}
}

func (srv *userServiceExecutor) RegisterKYC(
	ctx context.Context,
	kyc request.KYCRequestModel,
	newFiles []request.FileRequestModel) {

	userId, parserError := uuid.Parse(kyc.UserId)
	exceptions.PanicLogging(parserError)

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

	_, checkAccError := srv.CheckAccountExistByUserId(kyc.UserId)
	if checkAccError != nil {
		exceptions.PanicLogging(checkAccError)
	}

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

	reqVerification := request.AccVerificationRequestModel{
		StatusId:      2,
		TwoFactorAuth: true,
	}

	// update account verification status
	updateError := srv.UserRepo.UpdateAccountVerification(userId, reqVerification, true)
	if updateError != nil {
		exceptions.PanicLogging(updateError)
	}

	kycReview := entities.KYCReviewRequest{
		Id:             uuid.New(),
		UserRef:        kyc.UserId,
		PreRevStatus:   "awaitingReview",
		CreatedAt:      time.Now(),
		ReviewStatusId: 1,
	}

	reviewReqError := srv.BackofficeRepo.InsertKYCReviewRequest(kycReview)
	if reviewReqError != nil {
		exceptions.PanicLogging(reviewReqError)
	}

}

func (srv *userServiceExecutor) ChangeAccountStatus(ctx context.Context, userId string, statusId uint8) {

	id, parseError := uuid.Parse(userId)

	if parseError != nil {
		exceptions.PanicLogging(parseError)
	}

	_, checkAccError := srv.CheckAccountExistByUserId(userId)

	if checkAccError != nil {
		exceptions.PanicLogging(checkAccError)
	}

	reqVerification := request.AccVerificationRequestModel{
		StatusId: 3,
	}

	err := srv.UserRepo.UpdateAccountVerification(id, reqVerification, false)
	exceptions.PanicLogging(err)

}
