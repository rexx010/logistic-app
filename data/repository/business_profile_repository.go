package repository

import (
	"context"
	"errors"
	"logisticApp/config"
	"logisticApp/data/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessProfileRepository interface {
	Create(ctx context.Context, profile *models.BusinessProfile) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.BusinessProfile, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) (*models.BusinessProfile, error)
	Update(ctx context.Context, profile *models.BusinessProfile) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.UserStatusEnum) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type businessProfileRepository struct {
	db *gorm.DB
}

func NewBusinessProfileRepository() BusinessProfileRepository {
	return &businessProfileRepository{db: config.DB}
}

func (response *businessProfileRepository) Create(ctx context.Context, profile *models.BusinessProfile) error {
	return response.db.WithContext(ctx).Create(profile).Error
}

func (response *businessProfileRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.BusinessProfile, error) {
	var profile models.BusinessProfile
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

func (response *businessProfileRepository) FindByUserID(ctx context.Context, userID uuid.UUID) (*models.BusinessProfile, error) {
	var profile models.BusinessProfile
	err := response.db.WithContext(ctx).
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

func (response *businessProfileRepository) Update(ctx context.Context, profile *models.BusinessProfile) error {
	return response.db.WithContext(ctx).Save(profile).Error
}

func (response *businessProfileRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status models.UserStatusEnum) error {
	return response.db.WithContext(ctx).
		Model(&models.BusinessProfile{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": status}).Error
}

func (response *businessProfileRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return response.db.WithContext(ctx).Delete(&models.BusinessProfile{}, "id = ?", id).Error
}
