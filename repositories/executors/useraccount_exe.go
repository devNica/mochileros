package executors

import (
	"context"

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
			ProfileId: 4,
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
