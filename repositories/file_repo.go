package repositories

import (
	"context"

	"github.com/devNica/mochileros/entities"
	"github.com/google/uuid"
)

type FileRepository interface {
	InsertAssetByHotelId(ctx context.Context, newFile entities.File, hotelId uuid.UUID) error
}
