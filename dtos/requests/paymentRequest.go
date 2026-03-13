package requests

import "github.com/google/uuid"

type InitiatePaymentRequest struct {
	DeliveryID    uuid.UUID `json:"delivery_id" validate:"required,uuid4"`
	PaymentMethod string    `json:"payment_method" validate:"required,oneof=card bank_transfer ussd"`
}

type PaystackWebhookRequest struct {
	Event string              `json:"event"`
	Data  PaystackWebhookData `json:"data"`
}

type PaystackWebhookData struct {
	Reference string  `json:"reference" validate:"required"`
	Amount    float64 `json:"amount"    validate:"required,gt=0"`
	Status    string  `json:"status"    validate:"required"`
	Channel   string  `json:"channel"`
}
