package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type ClassRepository struct {
	DB *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{DB: db}
}

func (r *ClassRepository) Create(ctx context.Context, class *models.Class) error {
	return r.DB.WithContext(ctx).Create(class).Error
}

func (r *ClassRepository) Update(ctx context.Context, class *models.Class) error {
	return r.DB.WithContext(ctx).Save(class).Error
}

func (r *ClassRepository) Delete(ctx context.Context, classID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Class{}, classID).Error
}

func (r *ClassRepository) GetByID(ctx context.Context, classID uint) (*models.Class, error) {
	var class models.Class
	if err := r.DB.WithContext(ctx).First(&class, classID).Error; err != nil {
		return nil, err
	}
	return &class, nil
}
