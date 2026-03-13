package requests

import (
	"logisticApp/data/models"

	"github.com/google/uuid"
)

type CreateDeliveryRequest struct {
	PickupAddress  string  `json:"pickup_address" validate:"required,min=5"`
	DropoffAddress string  `json:"dropoff_address" validate:"required,min=5"`
	DeliveryFee    float64 `json:"delivery_fee" validate:"required,gt=0"`
}

type AssignRiderRequest struct {
	RiderID uuid.UUID `json:"rider_id" validate:"required,uuid4"`
}

type UpdateDeliveryStatusRequest struct {
	Status models.DeliveryStatusEnum `json:"status" validate:"required,oneof=IN_TRANSIT DELIVERED CANCELLED"`
}

type UploadProofRequest struct {
	ProofImageUrl string `json:"proof_image_url" validate:"required,url"`
}

type DeliveryFilterRequest struct {
	Status   models.DeliveryStatusEnum `form:"status"    validate:"omitempty,oneof=CREATED PAYMENT_PENDING PAID PAYMENT_FAILED ASSIGNED IN_TRANSIT DELIVERED COMPLETED CANCELLED"`
	Page     int                       `form:"page"      validate:"omitempty,min=1"`
	PageSize int                       `form:"page_size" validate:"omitempty,min=1,max=100"`
}
