package handlers

import (
	"net/http"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/helpers"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
)

type ClassHandler struct {
	ClassRepo *repositories.ClassRepository
}

func NewClassHandler(classRepo *repositories.ClassRepository) *ClassHandler {
	return &ClassHandler{ClassRepo: classRepo}
}

func (h *ClassHandler) CreateClass(c *fiber.Ctx) error {
	var class models.Class
	if err := c.BodyParser(&class); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.ClassRepo.Create(c.Context(), &class); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(class)
}

func (h *ClassHandler) UpdateClass(c *fiber.Ctx) error {
	classID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid class ID"})
	}

	var class models.Class
	if err := c.BodyParser(&class); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	class.ID = uint(classID)

	if err := h.ClassRepo.Update(c.Context(), &class); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(class)
}

func (h *ClassHandler) DeleteClass(c *fiber.Ctx) error {
	classID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid class ID"})
	}

	if err := h.ClassRepo.Delete(c.Context(), uint(classID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Class deleted successfully"})
}

func (h *ClassHandler) GetClassByID(c *fiber.Ctx) error {
	classID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid class ID"})
	}

	class, err := h.ClassRepo.GetByID(c.Context(), uint(classID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Class not found"})
	}

	return c.Status(http.StatusOK).JSON(class)
}
