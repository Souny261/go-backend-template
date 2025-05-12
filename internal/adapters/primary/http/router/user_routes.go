package router

import (
	handler "backend/internal/adapters/primary/http/handler"

	"github.com/gofiber/fiber/v2"
)

// RegisterAuthRoutes registers all authentication related routes
func SetupUserRoutes(router fiber.Router, userHandler handler.UserHandler) {
	router.Get("/users", userHandler.GetUser)
	router.Get("/users/:id", userHandler.GetUserByID)
	router.Put("/users", userHandler.UpdateUser)
	router.Delete("/users/:id", userHandler.DeleteUser)
	router.Post("/users/avatar", userHandler.UploadAvatar)
}
