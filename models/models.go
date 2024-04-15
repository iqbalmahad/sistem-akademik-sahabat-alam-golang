package models

import (
	"gorm.io/gorm"
)

// User represents the User entity
type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Username string `gorm:"uniqueIndex" json:"username" validate:"required"`
	Password string `json:"-" validate:"required,min=6"`
	Role     string `validate:"required,oneof=student teacher admin"`

	Student Student `validate:"-"`
	Teacher Teacher `validate:"-"`
	Admin   Admin   `validate:"-"`
}

// Student represents the Student entity
type Student struct {
	gorm.Model
	Name    string   `json:"name" validate:"required"`
	Nis     string   `gorm:"uniqueIndex" json:"nis" validate:"required"`
	Reports []Report `validate:"-"`
	ClassID *uint    `json:"class_id" validate:"-"`
	UserID  uint     `validate:"-"`
}

// Teacher represents the Teacher entity
type Teacher struct {
	gorm.Model
	Name    string `json:"name" validate:"required"`
	UserID  uint   `gorm:"unique" json:"-" validate:"-"`
	ClassID *uint  `json:"class_id" validate:"-"`
}

// Class represents the Class entity
type Class struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	SchoolID int    `json:"school_id" validate:"-"`

	Students []Student `validate:"-"`
	Teachers []Teacher `validate:"-"`
}

// School represents the School entity
type School struct {
	gorm.Model
	Name string `json:"name" validate:"required"`

	Classes []Class `validate:"-"`
}

// Admin represents the Admin entity
type Admin struct {
	gorm.Model
	Name   string `json:"name" validate:"required"`
	UserID uint   `validate:"-"`
}

// Report represents the Report entity
type Report struct {
	gorm.Model
	StudentID int    `json:"student_id" validate:"-"`
	Name      string `json:"name" validate:"required"`
}

// GraduationYearTK represents the GraduationYearTK entity
type GraduationYearTk struct {
	gorm.Model
	StudentID uint `json:"student_id" validate:"-"`
	Year      int  `json:"year"`

	Students []Student `validate:"-"`
}

// GraduationYearSD represents the GraduationYearSD entity
type GraduationYearSd struct {
	gorm.Model
	StudentID uint `json:"student_id" validate:"-"`
	Year      int  `json:"year"`

	Students []Student `validate:"-"`
}

// GraduationYearSMP represents the GraduationYearSMP entity
type GraduationYearSmp struct {
	gorm.Model
	StudentID uint `json:"student_id" validate:"-"`
	Year      int  `json:"year"`

	Students []Student `validate:"-"`
}
