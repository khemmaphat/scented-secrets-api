package infService

import (
	"context"

	"github.com/khemmaphat/scented-secrets-api/src/entities"
)

type IUserService interface {
	GetUserById(ctx context.Context, id string) (entities.User, error)
	CrateUser(ctx context.Context, user entities.User) error
	LoginUser(ctx context.Context, user entities.User) (string, error)
	EditUser(ctx context.Context, id string, user entities.User) error
}
