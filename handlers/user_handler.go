package handlers

import (
	"context"
	"net/http"
	"strconv"

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

	return c.Redirect("/users")
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserRepo.GetAll(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	data := fiber.Map{
		"Users": users,
	}
	// return c.Render("users/index", data)
	return c.Render("users/index", data, "layouts/main")
	// return c.Render("layouts/main", data)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid user ID")
	}

	user, err := h.UserRepo.GetByID(context.Background(), uint(idUint))
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	data := fiber.Map{
		"User": user,
	}
	// Render tampilan HTML dengan data pengguna
	return c.Render("users/show", data, "layouts/main")
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var updateUser models.User
	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Failed to parse request body")
	}

	// Validasi input menggunakan validator
	if err := validate.Struct(updateUser); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Hash password jika ada perubahan
	if updateUser.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Failed to hash password")
		}
		updateUser.Password = string(hashedPassword)
	}

	if err := h.UserRepo.Update(context.Background(), &updateUser); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect("/users")
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid user ID")
	}

	if err := h.UserRepo.Delete(context.Background(), uint(idUint)); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(http.StatusOK).SendString("User deleted successfully")
}
