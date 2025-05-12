package input

import (
	"backend/internal/core/domain"
	"context"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	// Manage User
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id uint) error
	UploadAvatar(ctx *fiber.Ctx, user *domain.User) error
}
