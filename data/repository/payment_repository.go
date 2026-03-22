package repository

import (
	"context"
	"errors"
	"logisticApp/config"
	"logisticApp/data/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *models.Payment) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Payment, error)
	FindByDeliveryID(ctx context.Context, deliveryID uuid.UUID) (*models.Payment, error)
	FindByReference(ctx context.Context, reference string) (*models.Payment, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.PaymentStatusEnum) error
	UpdateReference(ctx context.Context, id uuid.UUID, reference string) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository() PaymentRepository {
	return &paymentRepository{db: config.DB}
}

func (response *paymentRepository) Create(ctx context.Context, payment *models.Payment) error {
	return response.db.WithContext(ctx).Create(payment).Error
}

func (response *paymentRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Payment, error) {
	var payment models.Payment
	err := response.db.WithContext(ctx).First(&payment, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &payment, nil
}

func (response *paymentRepository) FindByDeliveryID(ctx context.Context, deliveryID uuid.UUID) (*models.Payment, error) {
	var payment models.Payment
	err := response.db.WithContext(ctx).
		Where("delivery_id = ?", deliveryID).
		First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &payment, nil
}

func (response *paymentRepository) FindByReference(ctx context.Context, reference string) (*models.Payment, error) {
	var payment models.Payment
	err := response.db.WithContext(ctx).
		Where("payment_reference = ?", reference).
		First(&payment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &payment, nil
}

func (response *paymentRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status models.PaymentStatusEnum) error {
	return response.db.WithContext(ctx).
		Model(&models.Payment{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": status}).Error
}

func (response *paymentRepository) UpdateReference(ctx context.Context, id uuid.UUID, reference string) error {
	return response.db.WithContext(ctx).
		Model(&models.Payment{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"payment_reference": reference}).Error
}
