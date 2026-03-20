package repository

import (
	"context"
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
