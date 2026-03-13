package responses

import (
	"logisticApp/data/models"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID             `json:"id"`
	Name      string                `json:"name"`
	Email     string                `json:"email"`
	Phone     string                `json:"phone"`
	Role      models.RoleEnum       `json:"role"`
	Status    models.UserStatusEnum `json:"status"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`

	BusinessProfile *BusinessProfileResponse `json:"business_profile,omitempty"`
	RiderProfile    *RiderProfileResponse    `json:"rider_profile,omitempty"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
	PaginationMeta
}
