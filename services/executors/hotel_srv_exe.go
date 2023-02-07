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
	repositories.FileRepository
}

func NewHotelServiceExecutor(hotelRepo *repositories.HotelRepo, fileRepo *repositories.FileRepository) services.HotelService {
	return &hotelServiceExecutor{HotelRepo: *hotelRepo, FileRepository: *fileRepo}
}

func (repo *hotelServiceExecutor) RegisterHotel(ctx context.Context, newHotel request.HotelRequestModel, newFile request.FileRequestModel) {
	commons.ValidateModel(newHotel)

	ownerId, errParse := uuid.Parse(newHotel.OwnerId)
	exceptions.PanicLogging(errParse)

	hotel := entities.Hotel{
		Id:                 uuid.New(),
		NameHotel:          newHotel.NameHotel,
		Address:            newHotel.Address,
		ServicePhoneNumber: newHotel.ServicePhoneNumber,
		State:              newHotel.State,
		Province:           newHotel.Province,
		CreatedAt:          time.Now(),
		OwnerId:            ownerId,
		CountryID:          newHotel.CountryID,
	}

	hotelId, err := repo.HotelRepo.InsertHotel(ctx, hotel)
	exceptions.PanicLogging(err)

	asset := entities.File{
		Filename: uuid.New(),
		Filetype: newFile.Filetype,
		Filesize: newFile.Filesize,
		Binary:   newFile.Buffer,
	}

	fileError := repo.FileRepository.InsertAssetByHotelId(ctx, asset, hotelId)
	exceptions.PanicLogging(fileError)

}

func (repo *hotelServiceExecutor) GetListOwnerHotels(ctx context.Context, ownerId string) []response.HotelResponseModel {
	hotelsE, error := repo.HotelRepo.FetchListOwnerHotels(ctx, ownerId)
	exceptions.PanicLogging(error)

	return hotelsE
}
