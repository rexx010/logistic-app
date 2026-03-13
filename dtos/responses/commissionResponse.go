package responses

import (
	"time"

	"github.com/google/uuid"
)

type CommissionResponse struct {
	ID          uuid.UUID `json:"id"`
	DeliveryID  uuid.UUID `json:"delivery_id"`
	PlatformFee float64   `json:"platform_fee"`
	RiderFee    float64   `json:"rider_fee"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
