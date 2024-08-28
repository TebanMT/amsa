package repository_impl

import (
	"context"

	"github.com/TebanMT/amsa/domain/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User
	if err := repo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) RegisterUser(ctx context.Context, user *entities.User) error {
	result := repo.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
