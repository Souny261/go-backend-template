package mysql

import (
	"backend/internal/adapters/secondary/mysql/exception"
	"backend/internal/core/domain"
	"backend/internal/core/ports/output"
	"context"
	"time"

	"gorm.io/gorm"
)

// userRepository implements the UserRepository interface using PostgreSQL
type UserRepository struct {
	db *gorm.DB
}

// VerifiedAccount implements output.UserRepository.
func (u *UserRepository) VerifiedAccount(ctx context.Context, user *domain.User) error {
	err := u.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":          user.Name,
			"password_hash": user.PasswordHash,
			"verified":      true,
			"status":        true,
		}).Error
	if err != nil {
		return exception.HandleGormError(err)
	}
	return nil
}

// FindByEmail implements output.UserRepository.
func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, exception.HandleGormError(err)
	}
	return &user, nil
}

// LastUserActive implements output.UserRepository.
func (u *UserRepository) LastUserActive(ctx context.Context, id uint) error {
	err := u.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", id).Update("last_active", time.Now()).Error
	if err != nil {
		return exception.HandleGormError(err)
	}
	return nil
}

// FindUserByUsernameAndPassword implements output.UserRepository.
func (u *UserRepository) FindUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	if err := u.db.WithContext(ctx).Where("username = ? OR email = ?", username, username).First(&user).Error; err != nil {
		return nil, exception.HandleGormError(err)
	}
	return &user, nil
}

// FindUserReleatedToRole implements output.UserRepository.
func (u *UserRepository) FindUserReleatedToRole(ctx context.Context, tenantID uint, roleID uint) uint {
	var count int64
	if err := u.db.
		WithContext(ctx).
		Model(&domain.User{}).
		Where("tenant_id = ?", tenantID).
		Where("role_id = ?", roleID).
		Count(&count).
		Error; err != nil {
		return 0
	}
	return uint(count)
}

// Delete implements output.UserRepository.
func (u *UserRepository) Delete(ctx context.Context, id uint) error {
	err := u.db.WithContext(ctx).Delete(&domain.User{}, id).Error
	if err != nil {
		return exception.HandleGormError(err)
	}
	return nil
}

// Update implements output.UserRepository.
func (u *UserRepository) Update(ctx context.Context, user *domain.User) error {
	err := u.db.WithContext(ctx).
		Where("id = ?", user.ID).
		Omit("Email", "Username", "Password").
		Save(user).Error
	if err != nil {
		return exception.HandleGormError(err)
	}
	return nil
}

// Create implements output.UserRepository.
func (u *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := u.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, exception.HandleGormError(err)
	}
	return user, nil
}

// GetAll implements output.UserRepository.
func (u *UserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := u.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, exception.HandleGormError(err)
	}
	return users, nil
}

// GetByID implements output.UserRepository.
func (u *UserRepository) FindByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := u.db.WithContext(ctx).
		Model(&domain.User{}).
		Preload("UserTenants").
		Preload("UserTenants.Tenant").
		Preload("UserTenants.Role").
		Preload("UserTenants.Role.RolePermissions").
		Preload("UserTenants.Role.RolePermissions.Permission").
		Preload("UserTenants.Role.RolePermissions.Service").
		Where("id = ?", id).
		First(&user).
		Error; err != nil {
		return nil, exception.HandleGormError(err)
	}
	return &user, nil
}

// NewUserRepository creates a new PostgreSQL user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

var _ output.UserRepository = (*UserRepository)(nil)
