package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, userID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.User{}, userID).Error
}

func (r *UserRepository) GetByID(ctx context.Context, userID uint) (*models.User, error) {
	var user models.User
	if err := r.DB.WithContext(ctx).First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	if err := r.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
