package repositories

import (
	"context"

	"github.com/devNica/mochileros/dto/response"
	"github.com/devNica/mochileros/entities"
	"github.com/google/uuid"
)

type HotelRepo interface {
	InsertHotel(ctx context.Context, entity entities.Hotel) (uuid.UUID, error)
	FetchListOwnerHotels(ctx context.Context, ownerId string) ([]response.HotelResponseModel, error)
}
