package repositories

import (
	"context"

	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
	"github.com/google/uuid"
)

type UserRepo interface {
	UserInsert(newUser entities.UserAccount, profileId uint16) error
	FetchUserByEmail(email string) (response.UserInfoResponseModel, error)
	InsertKYC(ctx context.Context, kyc entities.UserInfo) error
	CheckAccountExistByUserId(userId string) (entities.UserAccount, error)
	UpdateAccountVerification(userId uuid.UUID, req request.AccVerificationRequestModel, isFull bool) error
}
