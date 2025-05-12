package application

import "backend/internal/adapters/primary/http/handler"

// handler.AuthHandler

type AppHandler struct {
	auth handler.AuthHandler
	user handler.UserHandler
}

func SetupHandlers(svc *AppServices) *AppHandler {
	return &AppHandler{
		auth: handler.NewAuthHandler(svc.auth),
		user: handler.NewUserHandler(svc.user),
	}
}
