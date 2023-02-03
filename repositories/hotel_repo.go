package repositories

import (
	"context"

	"github.com/devNica/mochileros/entities"
)

type HotelRepo interface {
	InsertHotel(ctx context.Context, entity entities.Hotel) error
}
