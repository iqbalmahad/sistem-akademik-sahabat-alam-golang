package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type GraduationYearSdRepository struct {
	DB *gorm.DB
}

func NewGraduationYearSdRepository(db *gorm.DB) *GraduationYearSdRepository {
	return &GraduationYearSdRepository{DB: db}
}

func (r *GraduationYearSdRepository) Create(ctx context.Context, graduationYear *models.GraduationYearSd) error {
	return r.DB.WithContext(ctx).Create(graduationYear).Error
}

func (r *GraduationYearSdRepository) Update(ctx context.Context, graduationYear *models.GraduationYearSd) error {
	return r.DB.WithContext(ctx).Save(graduationYear).Error
}

func (r *GraduationYearSdRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.GraduationYearSd{}, id).Error
}

func (r *GraduationYearSdRepository) GetByID(ctx context.Context, id uint) (*models.GraduationYearSd, error) {
	var graduationYear models.GraduationYearSd
	if err := r.DB.WithContext(ctx).First(&graduationYear, id).Error; err != nil {
		return nil, err
	}
	return &graduationYear, nil
}
