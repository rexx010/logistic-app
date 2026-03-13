package requests

import "logisticApp/data/models"

type CreateRiderProfileRequest struct {
	VehicleType   models.VehicleTypeEnum `json:"vehicle_type"   validate:"required,oneof=BIKE CAR TRICYCLE VAN"`
	LicenseNumber string                 `json:"license_number" validate:"required,min=3,max=50"`
}

type UpdateRiderProfileRequest struct {
	VehicleType   models.VehicleTypeEnum `json:"vehicle_type"   validate:"omitempty,oneof=BIKE CAR TRICYCLE VAN"`
	LicenseNumber string                 `json:"license_number" validate:"omitempty,min=3,max=50"`
}
