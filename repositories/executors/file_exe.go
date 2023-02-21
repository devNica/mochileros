package executors

import (
	"context"
	"errors"

	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type fileRepositoryExecutor struct {
	*gorm.DB
}

func NewFileRepositoryExecutor(DB *gorm.DB) repositories.FileRepo {
	return &fileRepositoryExecutor{DB: DB}
}

func (repo *fileRepositoryExecutor) InsertAssetByHotelId(
	ctx context.Context,
	newFile entities.File,
	hotelId uuid.UUID) error {

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

func (repo *fileRepositoryExecutor) InsertAssetByUserId(
	ctx context.Context,
	newFile entities.File,
	assetTypeId uint16,
	userId uuid.UUID) error {

	err := repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newFile).Error; err != nil {
			tx.Rollback()
			return err
		}

		userAsset := entities.UserAssets{
			FileId:      newFile.Filename,
			UserId:      userId,
			AssetTypeId: assetTypeId,
		}

		if err := tx.Create(&userAsset).Error; err != nil {
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

func (repo *fileRepositoryExecutor) FetchHotelAsset(
	ctx context.Context,
	fileModel entities.File) (entities.File, error) {

	result := repo.DB.Select("binary").First(&fileModel)

	if result.RowsAffected == 0 {
		return entities.File{}, errors.New("Asset not found")
	}

	return fileModel, nil
}
