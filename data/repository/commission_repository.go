package repository

import (
	"context"
	"errors"
	"logisticApp/config"
	"logisticApp/data/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommissionRepository interface {
	Create(ctx context.Context, commission *models.Commission) error
	FindByDeliveryID(ctx context.Context, deliveryID uuid.UUID) (*models.Commission, error)
}

type commissionRepository struct {
	db *gorm.DB
}

func NewCommissionRepository() CommissionRepository {
	return &commissionRepository{db: config.DB}
}

func (response *commissionRepository) Create(ctx context.Context, commission *models.Commission) error {
	return response.db.WithContext(ctx).Create(commission).Error
}

func (response *commissionRepository) FindByDeliveryID(ctx context.Context, deliveryID uuid.UUID) (*models.Commission, error) {
	var commission models.Commission
	err := response.db.WithContext(ctx).
		Where("delivery_id = ?", deliveryID).
		First(&commission).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &commission, nil
}
