package repository

import (
	"context"
	"errors"
	"logisticApp/config"
	"logisticApp/data/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	Create(ctx context.Context, invoice *models.Invoice) error
	FindByID(ctx context.Context, id uuid.UUID) (*models.Invoice, error)
	FindByDeliveryID(ctx context.Context, deliveryID uuid.UUID) (*models.Invoice, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.InvoiceStatusEnum) error
	UpdatePDFURL(ctx context.Context, id uuid.UUID, url string) error
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{db: config.DB}
}

func (response *invoiceRepository) Create(ctx context.Context, invoice *models.Invoice) error {
	return response.db.WithContext(ctx).Create(invoice).Error
}

func (response *invoiceRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.Invoice, error) {
	var invoice models.Invoice
	err := response.db.WithContext(ctx).First(&invoice, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &invoice, nil
}

func (response *invoiceRepository) FindByDeliveryID(ctx context.Context, deliveryID uuid.UUID) (*models.Invoice, error) {
	var invoice models.Invoice
	err := response.db.WithContext(ctx).
		Where("delivery_id = ?", deliveryID).
		First(&invoice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &invoice, nil
}

func (response *invoiceRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status models.InvoiceStatusEnum) error {
	return response.db.WithContext(ctx).
		Model(&models.Invoice{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": status}).Error
}

func (response *invoiceRepository) UpdatePDFURL(ctx context.Context, id uuid.UUID, url string) error {
	return response.db.WithContext(ctx).
		Model(&models.Invoice{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"pdf_url": url}).Error
}
