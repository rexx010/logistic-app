package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationTypeEnum string

const (
	Email NotificationTypeEnum = "EMAIL"
	SMS   NotificationTypeEnum = "SMS"
	Push  NotificationTypeEnum = "PUSH"
)

type NotificationTriggerEnum string

const (
	PaymentVerified           NotificationTriggerEnum = "PAYMENT_VERIFIED"
	PaymentFailedNotification NotificationTriggerEnum = "PAYMENT_FAILED"
	RiderAssigned             NotificationTriggerEnum = "RIDER_ASSIGNED"
	InTransitNotification     NotificationTriggerEnum = "IN_TRANSIT"
	DeliveredNotification     NotificationTriggerEnum = "DELIVERED"
	RiderPaid                 NotificationTriggerEnum = "RIDER_PAID"
)

type NotificationStatusEnum string

const (
	NotificationStatusPending NotificationStatusEnum = "PENDING"
	NotificationStatusSent    NotificationStatusEnum = "SENT"
	NotificationStatusFailed  NotificationStatusEnum = "FAILED"
)

type Notification struct {
	Base
	UserID      uuid.UUID               `gorm:"type:uuid;not null"                 json:"user_id"`
	Type        NotificationTypeEnum    `gorm:"type:varchar(20)"                   json:"type"`
	Message     string                  `gorm:"not null"                            json:"message"`
	TriggeredBy NotificationTriggerEnum `gorm:"type:varchar(30)"                   json:"triggered_by"`
	Status      NotificationStatusEnum  `gorm:"type:varchar(20);default:'PENDING'" json:"status"`
	SentAt      *time.Time              `                                           json:"sent_at"`
	User        User                    `gorm:"foreignKey:UserID" json:"-"`
}
