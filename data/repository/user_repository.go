package repository

import (
	"context"
	"errors"
	"logisticApp/config"
	"logisticApp/data/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindById(ctx context.Context, id uuid.UUID) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByPhone(ctx context.Context, phone string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	UpdateProfilePicture(ctx context.Context, id uuid.UUID, url string) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.UserStatusEnum) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, filters UserFilters) ([]models.User, int64, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
}

type UserFilters struct {
	Role     models.RoleEnum
	Status   models.UserStatusEnum
	Page     int
	PageSize int
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{db: config.DB}
}

func (response *userRepository) Create(ctx context.Context, user *models.User) error {
	return response.db.WithContext(ctx).Create(user).Error
}

func (response *userRepository) FindById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	err := response.db.WithContext(ctx).
		Preload("BusinessProfile").
		Preload("RiderProfile").
		First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (response *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := response.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (response *userRepository) FindByPhone(ctx context.Context, phone string) (*models.User, error) {
	var user models.User
	err := response.db.WithContext(ctx).
		Where("phone = ?", phone).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (response *userRepository) Update(ctx context.Context, user *models.User) error {
	return response.db.WithContext(ctx).Save(user).Error
}

func (response *userRepository) UpdateProfilePicture(ctx context.Context, id uuid.UUID, url string) error {
	return response.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"profile_picture": url}).Error
}

func (response *userRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status models.UserStatusEnum) error {
	return response.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": status}).Error
}

func (response *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return response.db.WithContext(ctx).Delete(&models.User{}, "id = ?", id).Error
}

// List returns a paginated, filtered list of users.
// It builds the query dynamically — only filters that are non-empty are applied.
func (r *userRepository) List(ctx context.Context, filters UserFilters) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.db.WithContext(ctx).Model(&models.User{})

	// Conditionally add filters — empty string means "no filter applied"
	if filters.Role != "" {
		query = query.Where("role = ?", filters.Role)
	}
	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}

	// Count total matching records BEFORE applying pagination.
	// This gives us the number needed for total_pages calculation.
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (filters.Page - 1) * filters.PageSize
	err := query.
		Offset(offset).
		Limit(filters.PageSize).
		Order("created_at DESC").
		Find(&users).Error

	return users, total, err
}

func (response *userRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := response.db.WithContext(ctx).
		Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).Error
	return count > 0, err
}

func (response *userRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	var count int64
	err := response.db.WithContext(ctx).
		Model(&models.User{}).
		Where("phone = ?", phone).
		Count(&count).Error
	return count > 0, err
}
