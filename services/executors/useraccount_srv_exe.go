package executors

import (
	"context"
	"time"

	"github.com/devNica/mochileros/commons"
	argon2 "github.com/devNica/mochileros/commons/argon"
	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
	"github.com/google/uuid"
)

type userAccountServiceExecutor struct {
	repositories.UserAccountRepo
	configurations.Argon2Config
}

func NewUserAccountSrvExecutor(repo *repositories.UserAccountRepo, argon *configurations.Argon2Config) services.UserAccountService {
	return &userAccountServiceExecutor{UserAccountRepo: *repo, Argon2Config: *argon}
}

func (srv *userAccountServiceExecutor) UserAccountRegister(ctx context.Context, newUser request.UserAccounRequestModel) {
	commons.ValidateModel(newUser)
	hash := argon2.GeneratePassworHash(newUser.Password, &srv.Argon2Config)
	account := entities.UserAccount{
		Email:     newUser.Email,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := srv.UserAccountRepo.UserInsert(ctx, account)
	exceptions.PanicLogging(err)
}

func (srv *userAccountServiceExecutor) RegisterKYC(ctx context.Context, kyc request.KYCRequestModel) {

	userId, err := uuid.Parse(kyc.UserId)
	exceptions.PanicLogging(err)

	kycEntity := entities.UserInfo{
		FirstName: kyc.FirstName,
		LastName:  kyc.LastName,
		UserId:    userId,
	}

	srv.UserAccountRepo.InsertKYC(ctx, kycEntity)
}

func (srv *userAccountServiceExecutor) GetUserByEmail(ctx context.Context, email string) response.UserResponseModel {
	account, err := srv.UserAccountRepo.FetchUserByEmail(ctx, email)

	if err != nil {
		panic(exceptions.NotFoundError{
			Message: err.Error(),
		})
	}

	return account
}

func (srv *userAccountServiceExecutor) GetCompleteUserInfo(ctx context.Context, userId string) response.UserInfoResponseModel {
	user, err := srv.UserAccountRepo.FetchCompleteUserInfo(ctx, userId)
	if err != nil {
		panic(exceptions.NotFoundError{
			Message: err.Error(),
		})
	}

	return user
}

func (srv *userAccountServiceExecutor) ChangeAccountStatus(ctx context.Context, userId string) response.UserResponseModel {
	response, err := srv.UserAccountRepo.UpdateUserAccountStatus(ctx, userId)
	exceptions.PanicLogging(err)
	return response
}
