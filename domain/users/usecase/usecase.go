package users

import (
	"context"
	"gin-api-sample/domain/users/models"
)

type UserUsecaseI interface {
	CreateUser(ctx context.Context, req *models.User) error
	GetUser(ctx context.Context, req *string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, req *models.User) error
	DeleteData(ctx context.Context, req *string) error
}
