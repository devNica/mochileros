package executors

import (
	"context"
	"errors"
	"fmt"

	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/models"
	"github.com/devNica/mochileros/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userAccountExecutor struct {
	*gorm.DB
}

func NewUserAccountExecutor(DB *gorm.DB) repositories.UserAccountRepo {
	return &userAccountExecutor{DB: DB}
}

func (repo *userAccountExecutor) UserInsert(ctx context.Context, userAccount entities.UserAccount) error {
	userAccount.Id = uuid.New()
	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&userAccount).Error; err != nil {
			tx.Rollback()
			return err
		}

		profile := entities.UserHasProfile{
			UserId:    userAccount.Id,
			ProfileId: 2,
		}

		if err := tx.Create(&profile).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (repo *userAccountExecutor) InsertKYC(ctx context.Context, kyc entities.UserInfo) error {
	kyc.Id = uuid.New()
	err := repo.DB.WithContext(ctx).Create(&kyc).Error
	exceptions.PanicLogging(err)
	return nil
}

func (repo *userAccountExecutor) FetchUserByEmail(ctx context.Context, email string) (entities.UserAccount, error) {

	var foundUser entities.UserAccount

	result := repo.DB.WithContext(ctx).Where("email = ?", email).First(&foundUser)
	if result.RowsAffected == 0 {
		return entities.UserAccount{}, errors.New("user Not Found")
	}

	fmt.Println("user account", foundUser)

	return foundUser, nil
}

func (repo *userAccountExecutor) FetchCompleteUserInfo(ctx context.Context, userId string) (entities.UserAccount, error) {

	var queryModel models.CompleteUserRequestModel

	result := repo.DB.Raw("Select ua.id, ua.email, ui.first_name, ui.id as user_id, ui.last_name, ua.is_active from user_account ua inner join user_info ui on ui.user_id = ua.id").Scan(&queryModel)

	if result.RowsAffected == 0 {
		return entities.UserAccount{}, errors.New("user Not Found")
	}

	account := entities.UserAccount{
		Id:       queryModel.Id,
		Email:    queryModel.Email,
		IsActive: queryModel.IsActive,
		UserKYC: entities.UserInfo{
			Id:        queryModel.UserId,
			FirstName: queryModel.FirstName,
			LastName:  queryModel.LastName,
			UserId:    queryModel.Id,
		},
	}

	return account, nil

}
