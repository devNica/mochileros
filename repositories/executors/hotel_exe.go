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

func (repo *hotelRepoExecutor) InsertHotel(ctx context.Context, hotel entities.Hotel) (uuid.UUID, error) {
	hotel.Id = uuid.New()
	err := repo.DB.WithContext(ctx).Create(&hotel).Error
	exceptions.PanicLogging(err)
	return hotel.Id, err
}

func (repo *hotelRepoExecutor) FetchListOwnerHotels(ctx context.Context, ownerId string) ([]response.HotelResponseModel, error) {

	var hotels []response.HotelResponseModel

	repo.DB.
		Table("hotel").
		Select("hotel.id as hotel_id, hotel.name_hotel, hotel.address, hotel.service_phone_number, hotel.state, hotel.province, c.name as country").
		Joins("inner join country c on c.id = hotel.country_id").
		Where("hotel.owner_id = ?", ownerId).Scan(&hotels)

	return hotels, nil
}
