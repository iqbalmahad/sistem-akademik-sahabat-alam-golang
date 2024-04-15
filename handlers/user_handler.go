package handlers

import (
	"context"
	"net/http"

	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/helpers"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Validator instance
var validate = validator.New()

type UserHandler struct {
	UserRepo *repositories.UserRepository
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: userRepo}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validasi input menggunakan validator
	if err := validate.Struct(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Hash password sebelum disimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	if err := h.UserRepo.Create(context.Background(), &user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

//update

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	user.ID = uint(userID)

	// Validasi input menggunakan validator
	if err := validate.Struct(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Hash password sebelum update
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		user.Password = string(hashedPassword)
	}

	if err := h.UserRepo.Update(context.Background(), &user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(user)
}

// Tambahkan fungsi DeleteUser, GetUserByID, GetAllUsers sesuai kebutuhan

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := h.UserRepo.Delete(context.Background(), uint(userID)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	userID, err := helpers.ParseID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	user, err := h.UserRepo.GetByID(context.Background(), uint(userID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserRepo.GetAll(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get users"})
	}

	return c.JSON(users)
}
