package services

import (
	repository "DDD/app/Repository"
	"DDD/app/entity"
)

type UserService interface {
	GetUserById(Id uint) (entity.User, error)
	CreateUser(user entity.User) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetUserById(id uint) (entity.User, error) {
	return s.userRepo.FindByIdWithRolesAndPermission(id)
}

func (s *userService) CreateUser(user entity.User) error {
	return s.userRepo.Create(user)
}
