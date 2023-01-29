package executors

import (
	"context"
	"errors"

	"github.com/devNica/mochileros/entities"
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

func (repo *userAccountExecutor) FetchUserByEmail(ctx context.Context, email string) (entities.UserAccount, error) {

	var foundUser entities.UserAccount

	result := repo.DB.WithContext(ctx).Where("email = ?", email).First(&foundUser)
	if result.RowsAffected == 0 {
		return entities.UserAccount{}, errors.New("user Not Found")
	}

	return foundUser, nil
}
