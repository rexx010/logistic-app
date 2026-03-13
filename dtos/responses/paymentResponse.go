package responses

import (
	"logisticApp/data/models"
	"time"

	"github.com/google/uuid"
)

type PaymentInitiatedResponse struct {
	PaymentID        uuid.UUID `json:"payment_id"`
	AuthorizationURL string    `json:"authorization_url"`
	Reference        string    `json:"reference"`
	Amount           float64   `json:"amount"`
}

type PaymentResponse struct {
	ID               uuid.UUID                `json:"id"`
	DeliveryID       uuid.UUID                `json:"delivery_id"`
	Amount           float64                  `json:"amount"`
	Status           models.PaymentStatusEnum `json:"status"`
	PaymentReference string                   `json:"payment_reference"`
	PaymentMethod    string                   `json:"payment_method"`
	CreatedAt        time.Time                `json:"created_at"`
	UpdatedAt        time.Time                `json:"updated_at"`
}
