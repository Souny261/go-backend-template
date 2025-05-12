package services

import (
	"backend/internal/core/domain"
	"backend/internal/core/ports/input"
	"backend/internal/core/ports/output"
	"context"
)

type AuthServiceImpl struct {
	userRepo  output.UserRepository
	cacheRepo output.CacheRepository
}

// LastUserActive implements input.AuthService.
func (a *AuthServiceImpl) LastUserActive(ctx context.Context, id uint) error {
	return a.userRepo.LastUserActive(ctx, id)
}

// WhoAmI implements input.AuthService.
func (a *AuthServiceImpl) WhoAmI(ctx context.Context, userID uint) (*domain.User, error) {
	user, err := a.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Login implements input.AuthService.
func (a *AuthServiceImpl) Login(ctx context.Context, username string) (*domain.User, error) {
	return a.userRepo.FindUserByUsername(ctx, username)
}

func NewAuthService(
	userRepo output.UserRepository,
	cacheRepo output.CacheRepository,
) input.AuthService {
	return &AuthServiceImpl{
		userRepo:  userRepo,
		cacheRepo: cacheRepo,
	}
}
