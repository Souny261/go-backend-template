package router

import (
	handler "backend/internal/adapters/primary/http/handler"
	"backend/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

// RegisterAuthRoutes registers all authentication related routes
func SetupAuthRoutes(router fiber.Router, authHandler handler.AuthHandler) {
	// Public Router - routes that don't require authentication
	public := router.Group("/")
	public.Post("/login", authHandler.Login)
	public.Get("/jwt/refresh", jwt.AccessRefreshToken, authHandler.RefreshToken)

	// Private Router - routes that require authentication
	private := router.Group("/auth", jwt.AccessToken)
	private.Get("/me", authHandler.WhoAmI)
}
