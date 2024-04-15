package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) Create(ctx context.Context, admin *models.Admin) error {
	return r.DB.WithContext(ctx).Create(admin).Error
}

func (r *AdminRepository) Update(ctx context.Context, admin *models.Admin) error {
	return r.DB.WithContext(ctx).Save(admin).Error
}

func (r *AdminRepository) Delete(ctx context.Context, adminID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Admin{}, adminID).Error
}

func (r *AdminRepository) GetByID(ctx context.Context, adminID uint) (*models.Admin, error) {
	var admin models.Admin
	if err := r.DB.WithContext(ctx).First(&admin, adminID).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
