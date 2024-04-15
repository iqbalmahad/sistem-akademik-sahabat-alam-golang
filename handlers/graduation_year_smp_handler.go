package handlers

import (
	"net/http"
	"strconv"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
)

type GraduationYearSmpHandler struct {
	GraduationYearSmpRepo *repositories.GraduationYearSmpRepository
}

func NewGraduationYearSmpHandler(graduationYearSmpRepo *repositories.GraduationYearSmpRepository) *GraduationYearSmpHandler {
	return &GraduationYearSmpHandler{GraduationYearSmpRepo: graduationYearSmpRepo}
}

func (h *GraduationYearSmpHandler) CreateGraduationYearSmp(c *fiber.Ctx) error {
	var graduationYear models.GraduationYearSmp
	if err := c.BodyParser(&graduationYear); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.GraduationYearSmpRepo.Create(c.Context(), &graduationYear); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(graduationYear)
}

func (h *GraduationYearSmpHandler) UpdateGraduationYearSmp(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var graduationYear models.GraduationYearSmp
	if err := c.BodyParser(&graduationYear); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	graduationYear.ID = uint(id)

	if err := h.GraduationYearSmpRepo.Update(c.Context(), &graduationYear); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(graduationYear)
}

func (h *GraduationYearSmpHandler) DeleteGraduationYearSmp(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.GraduationYearSmpRepo.Delete(c.Context(), uint(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Graduation year deleted successfully"})
}

func (h *GraduationYearSmpHandler) GetGraduationYearSmpByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	graduationYear, err := h.GraduationYearSmpRepo.GetByID(c.Context(), uint(id))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Graduation year not found"})
	}

	return c.Status(http.StatusOK).JSON(graduationYear)
}
