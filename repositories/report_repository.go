package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (r *ReportRepository) Create(ctx context.Context, report *models.Report) error {
	return r.DB.WithContext(ctx).Create(report).Error
}

func (r *ReportRepository) Update(ctx context.Context, report *models.Report) error {
	return r.DB.WithContext(ctx).Save(report).Error
}

func (r *ReportRepository) Delete(ctx context.Context, reportID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Report{}, reportID).Error
}

func (r *ReportRepository) GetByID(ctx context.Context, reportID uint) (*models.Report, error) {
	var report models.Report
	if err := r.DB.WithContext(ctx).First(&report, reportID).Error; err != nil {
		return nil, err
	}
	return &report, nil
}
