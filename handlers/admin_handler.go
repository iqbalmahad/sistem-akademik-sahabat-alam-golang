package handlers

import (
	"net/http"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/helpers"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	AdminRepo *repositories.AdminRepository
}

func NewAdminHandler(adminRepo *repositories.AdminRepository) *AdminHandler {
	return &AdminHandler{AdminRepo: adminRepo}
}

func (h *AdminHandler) CreateAdmin(c *fiber.Ctx) error {
	var admin models.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.AdminRepo.Create(c.Context(), &admin); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(admin)
}

func (h *AdminHandler) UpdateAdmin(c *fiber.Ctx) error {
	adminID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid admin ID"})
	}

	var admin models.Admin
	if err := c.BodyParser(&admin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	admin.ID = uint(adminID)

	if err := h.AdminRepo.Update(c.Context(), &admin); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(admin)
}

func (h *AdminHandler) DeleteAdmin(c *fiber.Ctx) error {
	adminID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid admin ID"})
	}

	if err := h.AdminRepo.Delete(c.Context(), uint(adminID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Admin deleted successfully"})
}

func (h *AdminHandler) GetAdminByID(c *fiber.Ctx) error {
	adminID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid admin ID"})
	}

	admin, err := h.AdminRepo.GetByID(c.Context(), uint(adminID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
	}

	return c.Status(http.StatusOK).JSON(admin)
}
