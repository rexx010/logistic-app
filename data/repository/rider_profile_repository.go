package repository

import (
	"context"
	"errors"
	"logisticApp/config"
	"logisticApp/data/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RiderProfileRepository interface {
	Create(ctx context.Context, profile *models.RiderProfile) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.RiderProfile, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) (*models.RiderProfile, error)
	Update(ctx context.Context, profile *models.RiderProfile) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.UserStatusEnum) error
	UpdateEarningsBalance(ctx context.Context, id uuid.UUID, amount float64) error
	Delete(ctx context.Context, id uuid.UUID) error
	ListAvailable(ctx context.Context) ([]models.RiderProfile, error)
}

type riderProfileRepository struct {
	db *gorm.DB
}

func NewRiderProfileRepository() RiderProfileRepository {
	return &riderProfileRepository{db: config.DB}
}

func (response *riderProfileRepository) Create(ctx context.Context, profile *models.RiderProfile) error {
	return response.db.WithContext(ctx).Create(profile).Error
}

func (response *riderProfileRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.RiderProfile, error) {
	var profile models.RiderProfile
	err := response.db.WithContext(ctx).
		Preload("User").
		First(&profile, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (response *riderProfileRepository) FindByUserID(ctx context.Context, userID uuid.UUID) (*models.RiderProfile, error) {
	var profile models.RiderProfile
	err := response.db.WithContext(ctx).
		Preload("User").
		Where("user_id = ?", userID).
		First(&profile).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (response *riderProfileRepository) Update(ctx context.Context, profile *models.RiderProfile) error {
	return response.db.WithContext(ctx).Save(profile).Error
}

func (response *riderProfileRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status models.UserStatusEnum) error {
	return response.db.WithContext(ctx).
		Model(&models.RiderProfile{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": status}).Error
}

func (response *riderProfileRepository) UpdateEarningsBalance(ctx context.Context, id uuid.UUID, amount float64) error {
	return response.db.WithContext(ctx).
		Model(&models.RiderProfile{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"earnings_balance": gorm.Expr("earnings_balance + ?", amount),
		}).Error
}

func (response *riderProfileRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return response.db.WithContext(ctx).Delete(&models.RiderProfile{}, "id = ?", id).Error
}

func (response *riderProfileRepository) ListAvailable(ctx context.Context) ([]models.RiderProfile, error) {
	var riders []models.RiderProfile
	err := response.db.WithContext(ctx).
		Preload("User").
		Where("status = ?", models.Active).
		Find(&riders).Error
	return riders, err
}
