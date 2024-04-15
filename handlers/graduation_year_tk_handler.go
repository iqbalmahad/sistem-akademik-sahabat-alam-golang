package handlers

import (
	"net/http"
	"strconv"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
)

type GraduationYearTkHandler struct {
	GraduationYearTkRepo *repositories.GraduationYearTkRepository
}

func NewGraduationYearTkHandler(graduationYearTkRepo *repositories.GraduationYearTkRepository) *GraduationYearTkHandler {
	return &GraduationYearTkHandler{GraduationYearTkRepo: graduationYearTkRepo}
}

func (h *GraduationYearTkHandler) CreateGraduationYearTk(c *fiber.Ctx) error {
	var graduationYear models.GraduationYearTk
	if err := c.BodyParser(&graduationYear); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.GraduationYearTkRepo.Create(c.Context(), &graduationYear); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(graduationYear)
}

func (h *GraduationYearTkHandler) UpdateGraduationYearTk(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var graduationYear models.GraduationYearTk
	if err := c.BodyParser(&graduationYear); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	graduationYear.ID = uint(id)

	if err := h.GraduationYearTkRepo.Update(c.Context(), &graduationYear); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(graduationYear)
}

func (h *GraduationYearTkHandler) DeleteGraduationYearTk(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.GraduationYearTkRepo.Delete(c.Context(), uint(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Graduation year deleted successfully"})
}

func (h *GraduationYearTkHandler) GetGraduationYearTkByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	graduationYear, err := h.GraduationYearTkRepo.GetByID(c.Context(), uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Graduation year not found"})
	}

	return c.Status(http.StatusOK).JSON(graduationYear)
}
