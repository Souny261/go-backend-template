package output

import (
	"backend/internal/core/domain"
	"context"
)

type UserRepository interface {

	//  Manage User
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	FindByID(ctx context.Context, id uint) (*domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uint) error

	// Auth
	FindUserByUsername(ctx context.Context, username string) (*domain.User, error)
	LastUserActive(ctx context.Context, id uint) error

	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	VerifiedAccount(ctx context.Context, user *domain.User) error
}
