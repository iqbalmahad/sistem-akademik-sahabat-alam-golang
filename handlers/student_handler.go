package handlers

import (
	"net/http"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/helpers"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	StudentRepo *repositories.StudentRepository
}

func NewStudentHandler(studentRepo *repositories.StudentRepository) *StudentHandler {
	return &StudentHandler{StudentRepo: studentRepo}
}

func (h *StudentHandler) CreateStudent(c *fiber.Ctx) error {
	var student models.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.StudentRepo.Create(c.Context(), &student); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(student)
}

func (h *StudentHandler) UpdateStudent(c *fiber.Ctx) error {
	studentID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}

	var student models.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	student.ID = uint(studentID)

	if err := h.StudentRepo.Update(c.Context(), &student); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(student)
}

func (h *StudentHandler) DeleteStudent(c *fiber.Ctx) error {
	studentID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}

	if err := h.StudentRepo.Delete(c.Context(), uint(studentID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Student deleted successfully"})
}

func (h *StudentHandler) GetStudentByID(c *fiber.Ctx) error {
	studentID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}

	student, err := h.StudentRepo.GetByID(c.Context(), uint(studentID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Student not found"})
	}

	return c.Status(http.StatusOK).JSON(student)
}

func (h *StudentHandler) GetAllStudents(c *fiber.Ctx) error {
	ctx := c.Context()

	students, err := h.StudentRepo.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get students"})
	}

	return c.JSON(students)
}
