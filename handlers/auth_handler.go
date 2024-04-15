package handlers

import (
	"net/http"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthRepo *repositories.AuthRepository
}

func NewAuthHandler(authRepo *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{AuthRepo: authRepo}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	// Cari pengguna berdasarkan username
	user, err := h.AuthRepo.GetUserByUsername(c.Context(), req.Username)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Bandingkan password yang di-hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Return success message or token
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"id":       user.ID,
			"name":     user.Name,
			"username": user.Username,
			"role":     user.Role,
			// Tambahkan data lain yang ingin kamu kirimkan
		},
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	// Proses logout (jika diperlukan)
	return c.JSON(fiber.Map{"message": "Logout successful"})
}

func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	// Proses perubahan password (jika diperlukan)
	return c.JSON(fiber.Map{"message": "Change password successful"})
}
