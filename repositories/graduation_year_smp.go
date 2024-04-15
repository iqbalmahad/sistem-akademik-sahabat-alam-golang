package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type GraduationYearSmpRepository struct {
	DB *gorm.DB
}

func NewGraduationYearSmpRepository(db *gorm.DB) *GraduationYearSmpRepository {
	return &GraduationYearSmpRepository{DB: db}
}

func (r *GraduationYearSmpRepository) Create(ctx context.Context, graduationYear *models.GraduationYearSmp) error {
	return r.DB.WithContext(ctx).Create(graduationYear).Error
}

func (r *GraduationYearSmpRepository) Update(ctx context.Context, graduationYear *models.GraduationYearSmp) error {
	return r.DB.WithContext(ctx).Save(graduationYear).Error
}

func (r *GraduationYearSmpRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.WithContext(ctx).Delete(&models.GraduationYearSmp{}, id).Error
}

func (r *GraduationYearSmpRepository) GetByID(ctx context.Context, id uint) (*models.GraduationYearSmp, error) {
	var graduationYear models.GraduationYearSmp
	if err := r.DB.WithContext(ctx).First(&graduationYear, id).Error; err != nil {
		return nil, err
	}
	return &graduationYear, nil
}
