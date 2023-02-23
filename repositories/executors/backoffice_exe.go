package executors

import (
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/repositories"
	"gorm.io/gorm"
)

type backofficeExecutor struct {
	*gorm.DB
}

func NewBackofficeExecutor(DB *gorm.DB) repositories.BackofficeRepo {
	return &backofficeExecutor{DB: DB}
}

func (repo *backofficeExecutor) InsertKYCReviewRequest(kycReview entities.KYCReviewRequest) error {
	err := repo.DB.Create(&kycReview).Error

	if err != nil {
		return err
	}

	return nil
}
