package repositories

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
)

type UserAccountRepo interface {
	UserInsert(ctx context.Context, userAccountEntity entities.UserAccount) error
	InsertKYC(ctx context.Context, kyc entities.UserInfo) error
	UpdateUserAccountStatus(ctx context.Context, userId string) (entities.UserAccount, error)
	FetchUserByEmail(ctx context.Context, email string) (response.UserResponseModel, error)
	FetchCompleteUserInfo(ctx context.Context, userId string) (response.UserInfoResponseModel, error)
}
