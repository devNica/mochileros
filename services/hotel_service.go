package services

import (
	"context"

	"github.com/devNica/mochileros/models"
)

type HotelService interface {
	RegisterHotel(ctx context.Context, model models.HotelRequestModel)
}
