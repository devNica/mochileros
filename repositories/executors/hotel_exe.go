package executors

import (
	"context"

	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type hotelRepoExecutor struct {
	*gorm.DB
}

func NewHotelRepoExecutor(DB *gorm.DB) repositories.HotelRepo {
	return &hotelRepoExecutor{DB: DB}
}

func (repo *hotelRepoExecutor) InsertHotel(ctx context.Context, hotel entities.Hotel) error {
	hotel.Id = uuid.New()
	err := repo.DB.WithContext(ctx).Create(&hotel).Error
	exceptions.PanicLogging(err)
	return nil
}
