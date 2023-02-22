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

	//validate data request
	commons.ValidateModel(newCustomer)
	//generate hash password
	hash := argon2.GeneratePassworHash(newCustomer.Password, &srv.Argon2Config)
	// recover status key from dictionary
	accountStatus := commons.GetAccStatusDictionary()
	statusId := commons.GetAccStatusId("unverifiableIdentity", accountStatus)

	account := entities.UserAccount{
		Email:         newCustomer.Email,
		Password:      hash,
		PhoneNumber:   newCustomer.PhoneNumber,
		TwoFactorAuth: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		StatusId:      statusId,
	}

	profiles := commons.GetProfileDataDictionary()
	profileId := commons.GetProfileId("CUSTOMERS", profiles)

	err := srv.UserRepo.UserInsert(account, profileId)
	exceptions.PanicLogging(err)
}

func (srv *authServiceExecutor) UserLogin(ctx context.Context, user request.UserAccounRequestModel) response.LoginResponseModel {

	result, err := srv.UserRepo.FetchUserByEmail(user.Email)
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
		Id:            result.Id,
		Email:         result.Email,
		TwoFactorAuth: result.TwoFactorAuth,
		UserInfo:      result.UserInfo,
		CreatedAt:     result.CreatedAt,
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
