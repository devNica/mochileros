package repositories

import (
	"context"

	"github.com/devNica/mochileros/entities"
)

type UserAccountRepo interface {
	UserInsert(ctx context.Context, userAccountEntity entities.UserAccount) error
	InsertKYC(ctx context.Context, kyc entities.UserInfo) error
	UpdateUserAccountStatus(ctx context.Context, userId string) (entities.UserAccount, error)
	FetchUserByEmail(ctx context.Context, email string) (entities.UserAccount, error)
	FetchCompleteUserInfo(ctx context.Context, userId string) (entities.UserAccount, error)
}
