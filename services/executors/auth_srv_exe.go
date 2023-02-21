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
)

type authServiceExecutor struct {
	repositories.UserRepo
	configurations.Argon2Config
}

func NewAuthSrvExecutor(
	repo *repositories.UserRepo,
	argon *configurations.Argon2Config) services.AuthService {
	return &authServiceExecutor{UserRepo: *repo, Argon2Config: *argon}
}

func (srv *authServiceExecutor) CustomerRegister(
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

	err := srv.UserRepo.UserInsert(ctx, account, profileId)
	exceptions.PanicLogging(err)
}

func (srv *authServiceExecutor) UserLogin(ctx context.Context, user request.UserAccounRequestModel) response.LoginResponseModel {

	result, err := srv.UserRepo.FetchUserByEmail(ctx, user.Email)
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
