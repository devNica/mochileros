package services

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/models"
)

type HotelService interface {
	RegisterHotel(ctx context.Context, model models.HotelRequestModel)
	GetAllByOwnerId(ctx context.Context, ownerId string) []response.HotelResponseModel
}
