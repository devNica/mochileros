package executors

import (
	"context"
	"time"

	"github.com/devNica/mochileros/commons"
	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/repositories"
	"github.com/devNica/mochileros/services"
	"github.com/google/uuid"
)

type hotelServiceExecutor struct {
	repositories.HotelRepo
}

func NewHotelServiceExecutor(repo *repositories.HotelRepo) services.HotelService {
	return &hotelServiceExecutor{HotelRepo: *repo}
}

func (repo *hotelServiceExecutor) RegisterHotel(ctx context.Context, newHotel request.HotelRequestModel) {
	commons.ValidateModel(newHotel)

	ownerId, errParse := uuid.Parse(newHotel.OwnerId)
	exceptions.PanicLogging(errParse)

	hotel := entities.Hotel{
		Id:                 uuid.New(),
		NameHotel:          newHotel.NameHotel,
		Address:            newHotel.Address,
		ServicePhoneNumber: newHotel.ServicePhoneNumber,
		State:              newHotel.State,
		IsActive:           true,
		Province:           newHotel.Province,
		CreatedAt:          time.Now(),
		OwnerId:            ownerId,
		CountryID:          newHotel.CountryID,
	}

	err := repo.HotelRepo.InsertHotel(ctx, hotel)
	exceptions.PanicLogging(err)
}

func (repo *hotelServiceExecutor) GetListOwnerHotels(ctx context.Context, ownerId string) []response.HotelResponseModel {
	hotelsE, error := repo.HotelRepo.FetchListOwnerHotels(ctx, ownerId)
	exceptions.PanicLogging(error)

	return hotelsE
}
