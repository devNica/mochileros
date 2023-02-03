package services

import (
	"context"

	"github.com/devNica/mochileros/models"
)

type ResourcesService interface {
	GetAll(ctx context.Context) []models.Country
}
