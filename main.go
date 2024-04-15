package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/database"
	"github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/routes"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	engine := html.New("./templates", ".html")
	// Inisialisasi Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")

	// Panggil fungsi SetupRoutes dari package routes
	routes.SetupRoutes(app, db)
}
