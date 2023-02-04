package executors

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
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

func (repo *hotelRepoExecutor) FetchAllByOwnerID(ctx context.Context, ownerId string) ([]response.HotelResponseModel, error) {

	var hotels []response.HotelResponseModel

	repo.DB.WithContext(ctx).
		Table("hotel").
		Select("hotel.id, hotel.name_hotel, hotel.address, hotel.service_phone_number, hotel.state, hotel.province, c.name as country").
		Joins("join country c on c.id = hotel.country_id").
		Find(&hotels).Where("hotel.owner_id = ?", ownerId)

	return hotels, nil
}
