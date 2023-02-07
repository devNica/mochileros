package executors

import (
	"context"

	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type fileRepositoryExecutor struct {
	*gorm.DB
}

func NewFileRepositoryExecutor(DB *gorm.DB) repositories.FileRepository {
	return &fileRepositoryExecutor{DB: DB}
}

func (repo *fileRepositoryExecutor) InsertAssetByHotelId(ctx context.Context, newFile entities.File, hotelId uuid.UUID) error {

	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newFile).Error; err != nil {
			tx.Rollback()
			return err
		}

		hotelAsset := entities.HotelAssets{
			FileId:  newFile.Filename,
			HotelId: hotelId,
		}

		if err := tx.Create(&hotelAsset).Error; err != nil {
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
