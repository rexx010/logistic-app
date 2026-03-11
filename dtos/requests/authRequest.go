package requests

import "logisticApp/data/models"

type RegisterRequest struct {
	Name     string          `json:"name" validate:"required,min=2,max=100"`
	Email    string          `json:"email" validate:"required,email"`
	Phone    string          `json:"phone" validate:"required,min=11,max=14"`
	Password string          `json:"password" validate:"required,min=8,max=15"`
	Role     models.RoleEnum `json:"role"     validate:"required,oneof=BUSINESS_OWNER RIDER"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password"     validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"        validate:"required,email"`
	OTP         string `json:"otp"          validate:"required,len=6"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}
