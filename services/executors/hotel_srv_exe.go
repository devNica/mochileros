package executors

import (
	"context"
	"time"

	"github.com/devNica/mochileros/commons"
	"github.com/devNica/mochileros/entities"
	"github.com/devNica/mochileros/exceptions"
	"github.com/devNica/mochileros/models"
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

func (repo *hotelServiceExecutor) RegisterHotel(ctx context.Context, hotelReq models.HotelRequestModel) {
	commons.ValidateModel(hotelReq)

	ownerId, errParse := uuid.Parse(hotelReq.OwnerId)
	exceptions.PanicLogging(errParse)

	hotel := entities.Hotel{
		Id:                 uuid.New(),
		NameHotel:          hotelReq.NameHotel,
		Address:            hotelReq.Address,
		ServicePhoneNumber: hotelReq.ServicePhoneNumber,
		State:              hotelReq.State,
		IsActive:           true,
		Province:           hotelReq.Province,
		CreatedAt:          time.Now(),
		OwnerId:            ownerId,
		CountryID:          hotelReq.CountryID,
	}

	err := repo.HotelRepo.InsertHotel(ctx, hotel)
	exceptions.PanicLogging(err)
}
