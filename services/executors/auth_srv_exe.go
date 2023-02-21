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
	configurations.Config
}

func NewUserAccountSrvExecutor(
	repo *repositories.UserAccountRepo,
	argon *configurations.Argon2Config) services.AuthService {
	return &userAccountServiceExecutor{UserAccountRepo: *repo, Argon2Config: *argon}
}

func (srv *userAccountServiceExecutor) CustomerRegister(
	ctx context.Context,
	newCustomer request.UserAccounRequestModel) {
	commons.ValidateModel(newCustomer)
	hash := argon2.GeneratePassworHash(newCustomer.Password, &srv.Argon2Config)
	account := entities.UserAccount{
		Email:     newCustomer.Email,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	profiles := commons.GetProfileDataDictionary()
	profileId := commons.GetProfileId("CUSTOMERS", profiles)

	err := srv.UserAccountRepo.UserInsert(ctx, account, profileId)
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

func (srv *userAccountServiceExecutor) UserLogin(ctx context.Context, user request.UserAccounRequestModel) response.LoginResponseModel {

	result, err := srv.UserAccountRepo.FetchUserByEmail(ctx, user.Email)
	if err != nil {
		panic(exceptions.NotFoundError{
			Message: err.Error(),
		})
	}

	match, matchErr := argon2.ComparePasswordAndHash(user.Password, result.Password, &srv.Argon2Config)

	if match != true {
		panic(exceptions.BadReqquestError{
			Message: matchErr.Error(),
		})
	}

	login := response.LoginResponseModel{
		Id:        result.Id,
		Email:     result.Email,
		IsActive:  result.IsActive,
		UserInfo:  result.UserInfo,
		CreatedAt: result.CreatedAt,
	}

	profiles := make([]map[string]interface{}, 0)

	for _, val := range result.Profile {
		item := make(map[string]interface{})
		item["Role"] = val
		profiles = append(profiles, item)
	}

	login.Token = commons.GenerateToken(login.Id, profiles)

	return login
}

func (srv *userAccountServiceExecutor) ChangeAccountStatus(ctx context.Context, userId string) response.UserResponseModel {
	response, err := srv.UserAccountRepo.UpdateUserAccountStatus(ctx, userId)
	exceptions.PanicLogging(err)
	return response
}