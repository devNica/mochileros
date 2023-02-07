package services

import (
	"context"

	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
)

type HotelService interface {
	RegisterHotel(ctx context.Context, newHotel request.HotelRequestModel, newFile request.FileRequestModel)
	GetListOwnerHotels(ctx context.Context, ownerId string) []response.HotelResponseModel
}
