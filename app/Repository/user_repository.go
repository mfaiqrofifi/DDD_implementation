package repository

import (
	"DDD/app/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByIdWithRolesAndPermission(id uint) (entity.User, error)
	Create(user entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByIdWithRolesAndPermission(id uint) (entity.User, error) {
	var user entity.User
	err := r.db.Preload("Roles.Permissions").First(&user, id).Error
	return user, err
}

func (r *userRepository) Create(user entity.User) error {
	return r.db.Create(&user).Error
}
