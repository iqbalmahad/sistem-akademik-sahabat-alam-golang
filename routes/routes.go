package routes

import (
	handler "github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/handlers"
	repository "github.com/iqbalmahad/sistem-akademik-sahabat-alam-golang/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Inisialisasi repository
	userRepo := repository.NewUserRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	classRepo := repository.NewClassRepository(db)
	schoolRepo := repository.NewSchoolRepository(db)
	adminRepo := repository.NewAdminRepository(db)
	reportRepo := repository.NewReportRepository(db)
	graduationYearTkRepo := repository.NewGraduationYearTkRepository(db)
	graduationYearSdRepo := repository.NewGraduationYearSdRepository(db)
	graduationYearSmpRepo := repository.NewGraduationYearSmpRepository(db)
	authRepo := repository.NewAuthRepository(db)
	// Inisialisasi handler
	userHandler := handler.NewUserHandler(userRepo, app)
	studentHandler := handler.NewStudentHandler(studentRepo)
	teacherHandler := handler.NewTeacherHandler(teacherRepo)
	classHandler := handler.NewClassHandler(classRepo)
	schoolHandler := handler.NewSchoolHandler(schoolRepo)
	adminHandler := handler.NewAdminHandler(adminRepo)
	reportHandler := handler.NewReportHandler(reportRepo)
	graduationYearTkHandler := handler.NewGraduationYearTkHandler(graduationYearTkRepo)
	graduationYearSdHandler := handler.NewGraduationYearSdHandler(graduationYearSdRepo)
	graduationYearSmpHandler := handler.NewGraduationYearSmpHandler(graduationYearSmpRepo)
	authHandler := handler.NewAuthHandler(authRepo)

	// Routing HTTP
	app.Use(cors.New())

	app.Get("/users", userHandler.GetAllUsers)
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUserByID)
	// app.Put("/users/:id", userHandler.UpdateUser)
	// app.Delete("/users/:id", userHandler.DeleteUser)
	app.Get("/users/create", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("users/create", fiber.Map{
			"Title": "User Create",
		})
	})

	app.Get("/students", studentHandler.GetAllStudents)
	app.Post("/students", studentHandler.CreateStudent)
	app.Get("/students/:id", studentHandler.GetStudentByID)
	app.Put("/students/:id", studentHandler.UpdateStudent)
	app.Delete("/students/:id", studentHandler.DeleteStudent)

	app.Get("/teachers", teacherHandler.GetAllTeachers)
	app.Post("/teachers", teacherHandler.CreateTeacher)
	app.Get("/teachers/:id", teacherHandler.GetTeacherByID)
	app.Put("/teachers/:id", teacherHandler.UpdateTeacher)
	app.Delete("/teachers/:id", teacherHandler.DeleteTeacher)

	app.Post("/classes", classHandler.CreateClass)
	app.Get("/classes/:id", classHandler.GetClassByID)
	app.Put("/classes/:id", classHandler.UpdateClass)
	app.Delete("/classes/:id", classHandler.DeleteClass)

	app.Post("/schools", schoolHandler.CreateSchool)
	app.Get("/schools/:id", schoolHandler.GetSchoolByID)
	app.Put("/schools/:id", schoolHandler.UpdateSchool)
	app.Delete("/schools/:id", schoolHandler.DeleteSchool)

	app.Post("/admins", adminHandler.CreateAdmin)
	app.Get("/admins/:id", adminHandler.GetAdminByID)
	app.Put("/admins/:id", adminHandler.UpdateAdmin)
	app.Delete("/admins/:id", adminHandler.DeleteAdmin)

	app.Post("/reports", reportHandler.CreateReport)
	app.Get("/reports/:id", reportHandler.GetReportByID)
	app.Put("/reports/:id", reportHandler.UpdateReport)
	app.Delete("/reports/:id", reportHandler.DeleteReport)

	app.Post("/graduation-years-tk", graduationYearTkHandler.CreateGraduationYearTk)
	app.Get("/graduation-years-tk/:id", graduationYearTkHandler.GetGraduationYearTkByID)
	app.Put("/graduation-years-tk/:id", graduationYearTkHandler.UpdateGraduationYearTk)
	app.Delete("/graduation-years-tk/:id", graduationYearTkHandler.DeleteGraduationYearTk)

	app.Post("/graduation-years-sd", graduationYearSdHandler.CreateGraduationYearSd)
	app.Get("/graduation-years-sd/:id", graduationYearSdHandler.GetGraduationYearSdByID)
	app.Put("/graduation-years-sd/:id", graduationYearSdHandler.UpdateGraduationYearSd)
	app.Delete("/graduation-years-sd/:id", graduationYearSdHandler.DeleteGraduationYearSd)

	app.Post("/graduation-years-smp", graduationYearSmpHandler.CreateGraduationYearSmp)
	app.Get("/graduation-years-smp/:id", graduationYearSmpHandler.GetGraduationYearSmpByID)
	app.Put("/graduation-years-smp/:id", graduationYearSmpHandler.UpdateGraduationYearSmp)
	app.Delete("/graduation-years-smp/:id", graduationYearSmpHandler.DeleteGraduationYearSmp)

	// Login endpoint
	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)
	app.Post("/change-password", authHandler.ChangePassword)

	// Mulai server
	app.Listen(":8080")
}
