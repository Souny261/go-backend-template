package services

import (
	"backend/internal/core/domain"
	"backend/internal/core/ports/input"
	"backend/internal/core/ports/output"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserServiceImpl struct {
	userRepo  output.UserRepository
	cacheRepo output.CacheRepository
	minioRepo output.MinIORepository
}

// UploadAvatar implements input.UserService.
func (u *UserServiceImpl) UploadAvatar(ctx *fiber.Ctx, user *domain.User) error {
	multipartForm, err := ctx.MultipartForm()
	if err != nil {
		return err
	}
	if len(multipartForm.File["avatar"]) > 0 {
		fileHeader := multipartForm.File["avatar"][0]
		filename, err := u.minioRepo.UploadSingleFile(fileHeader, "images")
		if err != nil {
			return err
		}
		user.Avatar = filename
		fmt.Println(" UploadAvatar ser: ", user.Avatar)
	} else {
		fmt.Println("Avatar upload is optional, no file provided")
	}
	return nil
}

// DeleteUser implements input.UserService.
func (u *UserServiceImpl) DeleteUser(ctx context.Context, id uint) error {
	_, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	// Delete the user
	if err := u.userRepo.Delete(ctx, id); err != nil {
		return err
	}

	// Invalidate cache
	cacheKey := fmt.Sprintf("user:id:%d", id)
	u.cacheRepo.Delete(ctx, cacheKey)
	return nil
}

// UpdateUser implements input.UserService.
func (u *UserServiceImpl) UpdateUser(ctx context.Context, user *domain.User) error {
	_, err := u.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return err
	}
	// Update the user
	if err := u.userRepo.Update(ctx, user); err != nil {
		return err
	}
	// Invalidate cache
	cacheKey := fmt.Sprintf("user:id:%d", user.ID)
	u.cacheRepo.Delete(ctx, cacheKey)
	return nil
}

// GetUserByID implements input.UserService.
func (u *UserServiceImpl) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	cacheKey := fmt.Sprintf("user:id:%d", id)
	// Check Redis cache
	cachedUser, err := u.cacheRepo.Get(ctx, cacheKey)
	if err == nil && cachedUser != "" {
		// Deserialize JSON to User object
		var user domain.User
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			return &user, nil
		}
	}

	// Get from MySQL
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Serialize user to JSON and cache it
	userJSON, err := json.Marshal(user)
	if err == nil {
		u.cacheRepo.Set(ctx, cacheKey, string(userJSON), 30*time.Minute)
	}

	return user, nil
}

// GetUsers implements input.UserService.
func (u *UserServiceImpl) GetUsers(ctx context.Context) ([]domain.User, error) {
	fmt.Println("🚀 GetUsers ser")
	return u.userRepo.GetAll(ctx)
}

// NewUserService creates a new UserService
func NewUserService(
	userRepo output.UserRepository,
	cacheRepo output.CacheRepository,
	minioRepo output.MinIORepository,
) input.UserService {
	return &UserServiceImpl{
		userRepo:  userRepo,
		cacheRepo: cacheRepo,
		minioRepo: minioRepo,
	}
}
