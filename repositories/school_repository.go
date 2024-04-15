package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type SchoolRepository struct {
	DB *gorm.DB
}

func NewSchoolRepository(db *gorm.DB) *SchoolRepository {
	return &SchoolRepository{DB: db}
}

func (r *SchoolRepository) Create(ctx context.Context, school *models.School) error {
	return r.DB.WithContext(ctx).Create(school).Error
}

func (r *SchoolRepository) Update(ctx context.Context, school *models.School) error {
	return r.DB.WithContext(ctx).Save(school).Error
}

func (r *SchoolRepository) Delete(ctx context.Context, schoolID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.School{}, schoolID).Error
}

func (r *SchoolRepository) GetByID(ctx context.Context, schoolID uint) (*models.School, error) {
	var school models.School
	if err := r.DB.WithContext(ctx).First(&school, schoolID).Error; err != nil {
		return nil, err
	}
	return &school, nil
}
