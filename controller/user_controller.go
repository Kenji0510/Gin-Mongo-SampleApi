package controller

import (
	"context"
	"gin-api-sample/domain/users/models"
	userRepo "gin-api-sample/domain/users/repository"
	userUsercase "gin-api-sample/domain/users/usecase"
	"log"
)

type UserServiceImpl struct {
	userRepo userRepo.UserRepositoryI
	ctx      context.Context
}

func (u UserServiceImpl) CreateUser(ctx context.Context, req *models.User) error {
	if ctx == nil {
		ctx = context.Background()
	}
	err := u.userRepo.InsertData(ctx, req)
	if err != nil {
		return err
	}
	log.Println("Successfully Inserted Data User")

	return nil
}

func (u UserServiceImpl) GetUser(ctx context.Context, req *string) (*models.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := u.userRepo.GetData(ctx, req)
	if err != nil {
		log.Println("Failed to show data user with default log")
		return list, err
	}

	return list, nil
}

func (u UserServiceImpl) GetAll(ctx context.Context) ([]models.User, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	list, err := u.userRepo.GetAllData(ctx)
	if err != nil {
		log.Println("Failed to show data user with default log")
		return list, err
	}

	return list, err
}

func (u UserServiceImpl) UpdateUser(ctx context.Context, req *models.User) error {
	if ctx == nil {
		ctx = context.Background()
	}

	err := u.userRepo.UpdateData(ctx, req)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return nil
}

func (u UserServiceImpl) DeleteData(ctx context.Context, req *string) error {
	if ctx == nil {
		ctx = context.Background()
	}

	err := u.userRepo.DeleteData(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func NewUserUsecase(userRepo userRepo.UserRepositoryI, ctx context.Context) userUsercase.UserUsecaseI {
	return &UserServiceImpl{
		userRepo: userRepo,
		ctx:      ctx,
	}
}
