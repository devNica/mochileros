package repositories

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
)

type HotelRepo interface {
	InsertHotel(ctx context.Context, entity entities.Hotel) error
	FetchAllByOwnerID(ctx context.Context, ownerId string) ([]response.HotelResponseModel, error)
}
