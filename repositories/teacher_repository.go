package repositories

import (
	"context"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"

	"gorm.io/gorm"
)

type TeacherRepository struct {
	DB *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{DB: db}
}

func (r *TeacherRepository) Create(ctx context.Context, teacher *models.Teacher) error {
	return r.DB.WithContext(ctx).Create(teacher).Error
}

func (r *TeacherRepository) Update(ctx context.Context, teacher *models.Teacher) error {
	return r.DB.WithContext(ctx).Save(teacher).Error
}

func (r *TeacherRepository) Delete(ctx context.Context, teacherID uint) error {
	return r.DB.WithContext(ctx).Delete(&models.Teacher{}, teacherID).Error
}

func (r *TeacherRepository) GetByID(ctx context.Context, teacherID uint) (*models.Teacher, error) {
	var teacher models.Teacher
	if err := r.DB.WithContext(ctx).First(&teacher, teacherID).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *TeacherRepository) GetAll(ctx context.Context) ([]*models.Teacher, error) {
	var teachers []*models.Teacher
	if err := r.DB.WithContext(ctx).Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}
