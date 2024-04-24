package service

import (
	"context"

	"github.com/khemmaphat/scented-secrets-api/src/entities"
	infRepo "github.com/khemmaphat/scented-secrets-api/src/repository/infRepo"
	infServ "github.com/khemmaphat/scented-secrets-api/src/service/infService"
)

type UserService struct {
	userRepository infRepo.IUserRepository
}

func MakeUserService(userRepository infRepo.IUserRepository) infServ.IUserService {
	return &UserService{userRepository: userRepository}
}

func (r UserService) GetUserById(ctx context.Context, id string) (entities.User, error) {
	user, error := r.userRepository.GetUserById(ctx, id)
	return user, error
}

func (r UserService) CrateUser(ctx context.Context, user entities.User) error {
	err := r.userRepository.CrateUser(ctx, user)
	return err
}

func (r UserService) LoginUser(ctx context.Context, user entities.User) (string, error) {
	id, err := r.userRepository.LoginUser(ctx, user)
	return id, err
}

func (r UserService) EditUser(ctx context.Context, id string, user entities.User) error {
	err := r.userRepository.EditUser(ctx, id, user)
	return err
}

func (r UserService) UpdateNameUser(ctx context.Context, id string, name string) error {
	err := r.userRepository.UpdateNameUser(ctx, id, name)
	return err
}
