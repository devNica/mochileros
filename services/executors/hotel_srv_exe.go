package executors

import (
	"context"
	"fmt"
	"time"

	"github.com/devNica/mochileros/commons"
	"github.com/devNica/mochileros/commons/encryptors/cipher"
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
	repositories.FileRepo
	repositories.UserRepo
}

// type FileModel struct {
// 	Filename string
// 	Filetype string
// }

// func (f FileModel) StructToString() string {
// 	return fmt.Sprintf("%+v", f)
// }

func NewHotelServiceExecutor(
	hotelRepo *repositories.HotelRepo,
	fileRepo *repositories.FileRepo,
	userRepo *repositories.UserRepo) services.HotelService {
	return &hotelServiceExecutor{
		HotelRepo: *hotelRepo,
		FileRepo:  *fileRepo,
		UserRepo:  *userRepo,
	}
}

func (repo *hotelServiceExecutor) RegisterHotel(ctx context.Context, newHotel request.HotelRequestModel, newFile request.FileRequestModel) {

	user, checkError := repo.UserRepo.CheckAccountExistByUserId(newHotel.OwnerId)

	if checkError != nil {
		exceptions.PanicLogging(checkError)
	}

	if user.StatusId != 3 {
		exceptions.PanicLogging(exceptions.BadReqquestError{
			Message: "the kyc of the account is not approved",
		})
	}

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

	fileError := repo.FileRepo.InsertAssetByHotelId(ctx, asset, hotelId)
	exceptions.PanicLogging(fileError)

}

func (repo *hotelServiceExecutor) GetListOwnerHotels(ctx context.Context, baseURL, ownerId string) []response.HotelResponseModel {
	hotelRepModel, error := repo.HotelRepo.FetchListOwnerHotels(ctx, ownerId)
	exceptions.PanicLogging(error)

	var hotels []response.HotelResponseModel

	for _, hotel := range hotelRepModel {

		// file := FileModel{
		// 	Filename: hotel.Filename.String(),
		// 	Filetype: hotel.Filetype,
		// }

		// v := file.StructToString()
		v := fmt.Sprintf("%s,%s", hotel.Filename, hotel.Filetype)

		url, err := cipher.Encrypt(v)
		exceptions.PanicLogging(err)

		prefix := "/mochileros/v1/props/"

		hotels = append(hotels, response.HotelResponseModel{
			HotelId:            hotel.HotelId,
			NameHotel:          hotel.NameHotel,
			Address:            hotel.Address,
			ServicePhoneNumber: hotel.ServicePhoneNumber,
			Country:            hotel.Country,
			State:              hotel.State,
			Province:           hotel.Province,
			Url:                fmt.Sprintf("%s%s%s", baseURL, prefix, url),
		})
	}

	return hotels
}
