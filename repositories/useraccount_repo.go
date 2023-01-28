package repositories

import (
	"context"

	"github.com/devNica/mochileros/entities"
)

type UserAccountRepo interface {
	UserInsert(ctx context.Context, userAccountEntity entities.UserAccount) error
}
