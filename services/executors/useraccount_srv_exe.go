package executors

import (
	"context"
	"time"

	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
)

type userAccountServiceExecutor struct {
	repositories.UserAccountRepo
}

func NewUserAccountSrvExecutor(repo *repositories.UserAccountRepo) services.UserAccountService {
	return &userAccountServiceExecutor{UserAccountRepo: *repo}
}

func (srv *userAccountServiceExecutor) UserAccountRegister(ctx context.Context, requestModel models.UserAccounRequestModel) {
	account := entities.UserAccount{
		Email:     requestModel.Email,
		Password:  requestModel.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := srv.UserAccountRepo.UserInsert(ctx, account)
	exceptions.PanicLogging(err)
}
