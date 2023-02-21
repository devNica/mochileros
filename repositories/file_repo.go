package repositories

import (
	"context"

	"github.com/devNica/mochileros/entities"
	"github.com/google/uuid"
)

type FileRepo interface {
	InsertAssetByHotelId(ctx context.Context, newFile entities.File, hotelId uuid.UUID) error
	InsertAssetByUserId(ctx context.Context, newFile entities.File, assetTypeId uint16, userId uuid.UUID) error
	FetchHotelAsset(ctx context.Context, model entities.File) (entities.File, error)
}
