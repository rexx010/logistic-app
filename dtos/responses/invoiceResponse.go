package responses

import (
	"logisticApp/data/models"
	"time"

	"github.com/google/uuid"
)

type InvoiceResponse struct {
	ID         uuid.UUID                `json:"id"`
	DeliveryID uuid.UUID                `json:"delivery_id"`
	Amount     float64                  `json:"amount"`
	Status     models.InvoiceStatusEnum `json:"status"`
	PDFURL     string                   `json:"pdf_url"`
	CreatedAt  time.Time                `json:"created_at"`
	UpdatedAt  time.Time                `json:"updated_at"`
}
