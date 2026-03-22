package repository

import (
	"context"
	"errors"
	"logisticApp/config"
	"logisticApp/data/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(ctx context.Context, notification *models.Notification) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Notification, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.NotificationStatusEnum) error
	ListByUser(ctx context.Context, userID uuid.UUID, filters NotificationFilters) ([]models.Notification, int64, error)
}

type NotificationFilters struct {
	Status   models.NotificationStatusEnum
	Type     models.NotificationTypeEnum
	Page     int
	PageSize int
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository() NotificationRepository {
	return &notificationRepository{db: config.DB}
}

func (response *notificationRepository) Create(ctx context.Context, notification *models.Notification) error {
	return response.db.WithContext(ctx).Create(notification).Error
}

func (response *notificationRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Notification, error) {
	var notification models.Notification
	err := response.db.WithContext(ctx).First(&notification, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &notification, nil
}

func (response *notificationRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status models.NotificationStatusEnum) error {
	return response.db.WithContext(ctx).
		Model(&models.Notification{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": status}).Error
}

func (response *notificationRepository) ListByUser(ctx context.Context, userID uuid.UUID, filters NotificationFilters) ([]models.Notification, int64, error) {
	var notifications []models.Notification
	var total int64

	query := response.db.WithContext(ctx).
		Model(&models.Notification{}).
		Where("user_id = ?", userID)

	if filters.Status != "" {
		query = query.Where("status = ?", filters.Status)
	}
	if filters.Type != "" {
		query = query.Where("type = ?", filters.Type)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filters.Page - 1) * filters.PageSize
	err := query.
		Offset(offset).
		Limit(filters.PageSize).
		Order("created_at DESC").
		Find(&notifications).Error

	return notifications, total, err
}
