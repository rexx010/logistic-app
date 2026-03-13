package models

import "github.com/google/uuid"

type DeliveryStatusEnum string

const (
	Created        DeliveryStatusEnum = "CREATED"
	PaymentPending DeliveryStatusEnum = "PAYMENT_PENDING"
	Paid           DeliveryStatusEnum = "PAID"
	PaymentFailed  DeliveryStatusEnum = "PAYMENT_FAILED"
	Assigned       DeliveryStatusEnum = "ASSIGNED"
	InTransit      DeliveryStatusEnum = "IN_TRANSIT"
	Delivered      DeliveryStatusEnum = "DELIVERED"
	Completed      DeliveryStatusEnum = "COMPLETED"
	Cancelled      DeliveryStatusEnum = "CANCELLED"
)

type Delivery struct {
	Base
	BusinessID     uuid.UUID          `gorm:"type:uuid,not null" json:"business_Id"`
	RiderID        *uuid.UUID         `gorm:"type:uuid" json:"rider_id"`
	PickupAddress  string             `gorm:"not null" json:"pickup_address"`
	DropoffAddress string             `gorm:"not null" json:"dropoff_address"`
	DeliveryFee    float64            `gorm:"type:decimal(12,2);not null" json:"delivery_fee"`
	Status         DeliveryStatusEnum `gorm:"type:varchar(30);default:'CREATED'" json:"status"`
	ProofImageURL  string             `json:"proof_image_url"`
	Business       BusinessProfile    `gorm:"foreignKey:BusinessID" json:"-"`
	Rider          *RiderProfile      `gorm:"foreignKey:RiderID"    json:"rider,omitempty"`
	Payment        *Payment           `gorm:"foreignKey:DeliveryID" json:"payment,omitempty"`
	Invoice        *Invoice           `gorm:"foreignKey:DeliveryID" json:"invoice,omitempty"`
	Commission     *Commission        `gorm:"foreignKey:DeliveryID" json:"commission,omitempty"`
}
