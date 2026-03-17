package models

import "github.com/google/uuid"

type Commission struct {
	Base
	DeliveryID  uuid.UUID `gorm:"type:uuid;not null;uniqueIndex" json:"delivery_id"`
	PlatformFee float64   `gorm:"type:decimal(12,2)"             json:"platform_fee"`
	RiderFee    float64   `gorm:"type:decimal(12,2)"             json:"rider_fee"`

	Delivery Delivery `gorm:"foreignKey:DeliveryID" json:"-"`
}
