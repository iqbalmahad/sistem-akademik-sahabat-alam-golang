package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type GraduationYearTkRepository struct {
	DB *gorm.DB
}

func NewGraduationYearTkRepository(db *gorm.DB) *GraduationYearTkRepository {
	return &GraduationYearTkRepository{DB: db}
}

func (r *GraduationYearTkRepository) Create(ctx context.Context, graduationYear *models.GraduationYearTk) error {
	return r.DB.WithContext(ctx).Create(graduationYear).Error
}

func (r *GraduationYearTkRepository) Update(ctx context.Context, graduationYear *models.GraduationYearTk) error {
	return r.DB.WithContext(ctx).Save(graduationYear).Error
}

func (r *GraduationYearTkRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.GraduationYearTk{}, id).Error
}

func (r *GraduationYearTkRepository) GetByID(ctx context.Context, id uint) (*models.GraduationYearTk, error) {
	var graduationYear models.GraduationYearTk
	if err := r.DB.WithContext(ctx).First(&graduationYear, id).Error; err != nil {
		return nil, err
	}
	return &graduationYear, nil
}
