package input

import (
	"backend/internal/core/domain"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, username string) (*domain.User, error)
	WhoAmI(ctx context.Context, userID uint) (*domain.User, error)
	LastUserActive(ctx context.Context, id uint) error
}
