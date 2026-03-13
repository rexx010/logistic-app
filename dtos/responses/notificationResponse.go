package responses

import (
	"logisticApp/data/models"
	"time"

	"github.com/google/uuid"
)

type NotificationResponse struct {
	ID          uuid.UUID                      `json:"id"`
	UserID      uuid.UUID                      `json:"user_id"`
	Type        models.NotificationTypeEnum    `json:"type"`
	Message     string                         `json:"message"`
	TriggeredBy models.NotificationTriggerEnum `json:"triggered_by"`
	Status      models.NotificationStatusEnum  `json:"status"`
	SentAt      *time.Time                     `json:"sent_at"`
	CreatedAt   time.Time                      `json:"created_at"`
}

type NotificationListResponse struct {
	Notifications []NotificationResponse `json:"notifications"`
	PaginationMeta
}
