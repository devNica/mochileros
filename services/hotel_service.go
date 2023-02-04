package services

import (
	"context"

	"github.com/devNica/mochileros/dto/request"
	"github.com/devNica/mochileros/dto/response"
)

type HotelService interface {
	RegisterHotel(ctx context.Context, newHotel request.HotelRequestModel)
	GetAllByOwnerId(ctx context.Context, ownerId string) []response.HotelResponseModel
}