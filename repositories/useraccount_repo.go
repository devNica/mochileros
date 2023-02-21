package repositories

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
)

type UserAccountRepo interface {
	UserInsert(ctx context.Context, newUser entities.UserAccount, profileId uint16) error
	FetchUserByEmail(ctx context.Context, email string) (response.UserInfoResponseModel, error)
	InsertKYC(ctx context.Context, kyc entities.UserInfo) error
	UpdateUserAccountStatus(ctx context.Context, userId string) (response.UserResponseModel, error)
}
