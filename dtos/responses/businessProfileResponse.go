package responses

import (
	"logisticApp/data/models"
	"time"

	"github.com/google/uuid"
)

type BusinessProfileResponse struct {
	ID               uuid.UUID             `json:"id"`
	UserID           uuid.UUID             `json:"user_id"`
	BusinessName     string                `json:"business_name"`
	BusinessAddress  string                `json:"business_address"`
	SubscriptionPlan string                `json:"subscription_plan"`
	Status           models.UserStatusEnum `json:"status"`
	CreatedAt        time.Time             `json:"created_at"`
	UpdatedAt        time.Time             `json:"updated_at"`
}
