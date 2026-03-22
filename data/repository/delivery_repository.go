package repository

import (
	"context"
	"errors"
	"logisticApp/config"
	"logisticApp/data/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeliveryRepository interface {
	Create(ctx context.Context, delivery *models.Delivery) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Delivery, error)
	Update(ctx context.Context, delivery *models.Delivery) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.DeliveryStatusEnum) error
	UpdateRider(ctx context.Context, id uuid.UUID, riderID uuid.UUID) error
	UpdateProofImage(ctx context.Context, id uuid.UUID, url string) error
	Delete(ctx context.Context, id uuid.UUID) error
	ListByBusiness(ctx context.Context, businessID uuid.UUID, filters DeliveryFilters) ([]models.Delivery, int64, error)
	ListByRider(ctx context.Context, riderID uuid.UUID, filters DeliveryFilters) ([]models.Delivery, int64, error)
	ListAll(ctx context.Context, filters DeliveryFilters) ([]models.Delivery, int64, error)
}

type DeliveryFilters struct {
	Status   models.DeliveryStatusEnum
	Page     int
	PageSize int
}

type deliveryRepository struct {
	db *gorm.DB
}

func NewDeliveryRepository() DeliveryRepository {
	return &deliveryRepository{db: config.DB}
}

func (response *deliveryRepository) Create(ctx context.Context, delivery *models.Delivery) error {
	return response.db.WithContext(ctx).Create(delivery).Error
}

func (response *deliveryRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Delivery, error) {
	var delivery models.Delivery
	err := response.db.WithContext(ctx).
		Preload("Rider").
		Preload("Rider.User").
		Preload("Payment").
		Preload("Invoice").
		Preload("Commission").
		First(&delivery, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &delivery, nil
}

func (response *deliveryRepository) Update(ctx context.Context, delivery *models.Delivery) error {
	return response.db.WithContext(ctx).Save(delivery).Error
}

func (response *deliveryRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status models.DeliveryStatusEnum) error {
	return response.db.WithContext(ctx).
		Model(&models.Delivery{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": status}).Error
}

func (response *deliveryRepository) UpdateRider(ctx context.Context, id uuid.UUID, riderID uuid.UUID) error {
	return response.db.WithContext(ctx).
		Model(&models.Delivery{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"rider_id": riderID,
			"status":   models.Assigned,
		}).Error
}

func (response *deliveryRepository) UpdateProofImage(ctx context.Context, id uuid.UUID, url string) error {
	return response.db.WithContext(ctx).
		Model(&models.Delivery{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"proof_image": url}).Error
}

func (response *deliveryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return response.db.WithContext(ctx).Delete(&models.Delivery{}, "id = ?", id).Error
}

func (response *deliveryRepository) ListByBusiness(ctx context.Context, businessID uuid.UUID, filters DeliveryFilters) ([]models.Delivery, int64, error) {
	return response.list(ctx, "business_id = ?", businessID, filters)
}

func (response *deliveryRepository) ListByRider(ctx context.Context, riderID uuid.UUID, filters DeliveryFilters) ([]models.Delivery, int64, error) {
	return response.list(ctx, "rider_id = ?", riderID, filters)
}

func (response *deliveryRepository) ListAll(ctx context.Context, filters DeliveryFilters) ([]models.Delivery, int64, error) {
	return response.list(ctx, "", nil, filters)
}

func (response *deliveryRepository) list(ctx context.Context, condition string, value interface{}, filters DeliveryFilters) ([]models.Delivery, int64, error) {
	var deliveries []models.Delivery
	var total int64

	query := response.db.WithContext(ctx).Model(&models.Delivery{})

	if condition != "" {
		query = query.Where(condition, value)
	}
	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filters.Page - 1) * filters.PageSize
	err := query.
		Preload("Rider").
		Preload("Rider.User").
		Offset(offset).
		Limit(filters.PageSize).
		Order("created_at DESC").
		Find(&deliveries).Error

	return deliveries, total, err
}
