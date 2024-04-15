package handlers

import (
	"net/http"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/helpers"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
)

type SchoolHandler struct {
	SchoolRepo *repositories.SchoolRepository
}

func NewSchoolHandler(schoolRepo *repositories.SchoolRepository) *SchoolHandler {
	return &SchoolHandler{SchoolRepo: schoolRepo}
}

func (h *SchoolHandler) CreateSchool(c *fiber.Ctx) error {
	var school models.School
	if err := c.BodyParser(&school); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.SchoolRepo.Create(c.Context(), &school); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(school)
}

func (h *SchoolHandler) UpdateSchool(c *fiber.Ctx) error {
	schoolID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid school ID"})
	}

	var school models.School
	if err := c.BodyParser(&school); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	school.ID = uint(schoolID)

	if err := h.SchoolRepo.Update(c.Context(), &school); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(school)
}

func (h *SchoolHandler) DeleteSchool(c *fiber.Ctx) error {
	schoolID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid school ID"})
	}

	if err := h.SchoolRepo.Delete(c.Context(), uint(schoolID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "School deleted successfully"})
}

func (h *SchoolHandler) GetSchoolByID(c *fiber.Ctx) error {
	schoolID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid school ID"})
	}

	school, err := h.SchoolRepo.GetByID(c.Context(), uint(schoolID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "School not found"})
	}

	return c.Status(http.StatusOK).JSON(school)
}
