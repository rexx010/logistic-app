package requests

import (
	"logisticApp/data/models"

	"github.com/google/uuid"
)

type NotificationFilterRequest struct {
	Status   models.NotificationStatusEnum `form:"status"    validate:"omitempty,oneof=PENDING SENT FAILED"`
	Type     models.NotificationTypeEnum   `form:"type"      validate:"omitempty,oneof=EMAIL SMS PUSH"`
	Page     int                           `form:"page"      validate:"omitempty,min=1"`
	PageSize int                           `form:"page_size" validate:"omitempty,min=1,max=100"`
}

type SendNotificationRequest struct {
	UserID  uuid.UUID                   `json:"user_id" validate:"required,uuid4"`
	Type    models.NotificationTypeEnum `json:"type"    validate:"required,oneof=EMAIL SMS PUSH"`
	Message string                      `json:"message" validate:"required,min=5,max=500"`
}
