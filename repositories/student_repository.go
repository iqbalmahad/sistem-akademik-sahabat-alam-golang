package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

func (r *StudentRepository) Create(ctx context.Context, student *models.Student) error {
	return r.DB.WithContext(ctx).Create(student).Error
}

func (r *StudentRepository) Update(ctx context.Context, student *models.Student) error {
	return r.DB.WithContext(ctx).Save(student).Error
}

func (r *StudentRepository) Delete(ctx context.Context, studentID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Student{}, studentID).Error
}

func (r *StudentRepository) GetByID(ctx context.Context, studentID uint) (*models.Student, error) {
	var student models.Student
	if err := r.DB.WithContext(ctx).First(&student, studentID).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) GetAll(ctx context.Context) ([]*models.Student, error) {
	var students []*models.Student
	if err := r.DB.WithContext(ctx).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
