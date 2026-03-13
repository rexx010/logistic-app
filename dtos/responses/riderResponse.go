package responses

import (
	"logisticApp/data/models"
	"time"

	"github.com/google/uuid"
)

type RiderProfileResponse struct {
	ID              uuid.UUID              `json:"id"`
	UserID          uuid.UUID              `json:"user_id"`
	VehicleType     models.VehicleTypeEnum `json:"vehicle_type"`
	LicenseNumber   string                 `json:"license_number"`
	EarningsBalance float64                `json:"earnings_balance"`
	Status          models.UserStatusEnum  `json:"status"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

type RiderPublicResponse struct {
	ID            uuid.UUID              `json:"id"`
	UserID        uuid.UUID              `json:"user_id"`
	Name          string                 `json:"name"`
	Phone         string                 `json:"phone"`
	VehicleType   models.VehicleTypeEnum `json:"vehicle_type"`
	LicenseNumber string                 `json:"license_number"`
	Status        models.UserStatusEnum  `json:"status"`
}
