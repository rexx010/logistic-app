package models

import "github.com/google/uuid"

type InvoiceStatusEnum string

const (
	InvoiceStatusUnpaid InvoiceStatusEnum = "UNPAID"
	InvoiceStatusPaid   InvoiceStatusEnum = "PAID"
	InvoiceStatusVoided InvoiceStatusEnum = "VOIDED"
)

type Invoice struct {
	Base
	DeliveryID uuid.UUID         `gorm:"type:uuid;not null;uniqueIndex"    json:"delivery_id"`
	Amount     float64           `gorm:"type:decimal(12,2);not null"       json:"amount"`
	Status     InvoiceStatusEnum `gorm:"type:varchar(20);default:'UNPAID'" json:"status"`
	PDFURL     string            `                                          json:"pdf_url"`
	Delivery   Delivery          `gorm:"foreignKey:DeliveryID" json:"-"`
}
