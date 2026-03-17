package requests

type CreateBusinessProfileRequest struct {
	BusinessName    string `json:"business_name"    validate:"required,min=2,max=150"`
	BusinessAddress string `json:"business_address" validate:"required,min=5"`
}

type UpdateBusinessProfileRequest struct {
	BusinessName     string `json:"business_name"     validate:"omitempty,min=2,max=150"`
	BusinessAddress  string `json:"business_address"  validate:"omitempty,min=5"`
	SubscriptionPlan string `json:"subscription_plan" validate:"omitempty,oneof=free basic premium enterprise"`
}
