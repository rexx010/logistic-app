package responses

import (
	"logisticApp/data/models"
	"time"

	"github.com/google/uuid"
)

type DeliveryResponse struct {
	ID             uuid.UUID                 `json:"id"`
	BusinessID     uuid.UUID                 `json:"business_id"`
	RiderID        *uuid.UUID                `json:"rider_id"`
	PickupAddress  string                    `json:"pickup_address"`
	DropoffAddress string                    `json:"dropoff_address"`
	DeliveryFee    float64                   `json:"delivery_fee"`
	Status         models.DeliveryStatusEnum `json:"status"`
	ProofImageURL  string                    `json:"proof_image_url,omitempty"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`

	Rider      *RiderPublicResponse `json:"rider,omitempty"`
	Payment    *PaymentResponse     `json:"payment,omitempty"`
	Invoice    *InvoiceResponse     `json:"invoice,omitempty"`
	Commission *CommissionResponse  `json:"commission,omitempty"`
}

type DeliveryListResponse struct {
	Deliveries []DeliveryResponse `json:"deliveries"`
	PaginationMeta
}
