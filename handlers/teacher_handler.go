package handlers

import (
	"net/http"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/helpers"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
)

type TeacherHandler struct {
	TeacherRepo *repositories.TeacherRepository
}

func NewTeacherHandler(teacherRepo *repositories.TeacherRepository) *TeacherHandler {
	return &TeacherHandler{TeacherRepo: teacherRepo}
}

func (h *TeacherHandler) CreateTeacher(c *fiber.Ctx) error {
	var teacher models.Teacher
	if err := c.BodyParser(&teacher); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.TeacherRepo.Create(c.Context(), &teacher); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(teacher)
}

func (h *TeacherHandler) UpdateTeacher(c *fiber.Ctx) error {
	teacherID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid teacher ID"})
	}

	var teacher models.Teacher
	if err := c.BodyParser(&teacher); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	teacher.ID = uint(teacherID)

	if err := h.TeacherRepo.Update(c.Context(), &teacher); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(teacher)
}

func (h *TeacherHandler) DeleteTeacher(c *fiber.Ctx) error {
	teacherID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid teacher ID"})
	}

	if err := h.TeacherRepo.Delete(c.Context(), uint(teacherID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Teacher deleted successfully"})
}

func (h *TeacherHandler) GetTeacherByID(c *fiber.Ctx) error {
	teacherID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid teacher ID"})
	}

	teacher, err := h.TeacherRepo.GetByID(c.Context(), uint(teacherID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Teacher not found"})
	}

	return c.Status(http.StatusOK).JSON(teacher)
}

func (h *TeacherHandler) GetAllTeachers(c *fiber.Ctx) error {
	ctx := c.Context()

	teachers, err := h.TeacherRepo.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get teachers"})
	}

	return c.JSON(teachers)
}
