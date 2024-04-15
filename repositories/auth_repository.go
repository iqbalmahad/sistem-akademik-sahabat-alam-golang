package repositories

import (
	"context"
	"errors"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
