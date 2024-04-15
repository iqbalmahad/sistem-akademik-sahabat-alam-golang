package handlers

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/models"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

type UserHandler struct {
	UserRepo *repositories.UserRepository
	App      *fiber.App
}

func NewUserHandler(userRepo *repositories.UserRepository, app *fiber.App) *UserHandler {
	return &UserHandler{UserRepo: userRepo, App: app}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to parse request body")
	}

	// Validasi input menggunakan validator
	if err := validate.Struct(user); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Hash password sebelum disimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Failed to hash password")
	}
	user.Password = string(hashedPassword)

	if err := h.UserRepo.Create(context.Background(), &user); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusCreated).Render("templates/users/user_created.html", fiber.Map{"User": user})
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserRepo.GetAll(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Menggunakan template HTML dengan nama index.gohtml di folder templates/users
	return c.Render("users/index", fiber.Map{
		"Users": users,
	})
}
