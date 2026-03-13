package models

import "github.com/google/uuid"

type PaymentStatusEnum string

const (
	PendingPayment PaymentStatusEnum = "PENDING"
	Success        PaymentStatusEnum = "SUCCESS"
	Failed         PaymentStatusEnum = "FAILED"
	Released       PaymentStatusEnum = "RELEASED"
	Refunded       PaymentStatusEnum = "REFUNDED"
)

type Payment struct {
	Base
	DeliveryID       uuid.UUID         `gorm:"type:uuid;not null;uniqueIndex"     json:"delivery_id"`
	Amount           float64           `gorm:"type:decimal(12,2);not null"        json:"amount"`
	Status           PaymentStatusEnum `gorm:"type:varchar(20);default:'PENDING'" json:"status"`
	PaymentReference string            `                                           json:"payment_reference"`
	PaymentMethod    string            `                                           json:"payment_method"`

	Delivery Delivery `gorm:"foreignKey:DeliveryID" json:"-"`
}
