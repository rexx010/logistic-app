package requests

import "logisticApp/data/models"

type UpdateUserRequest struct {
	Name  string `json:"name"   validate:"omitempty,min=2,max=100"`
	Phone string `json:"phone" validate:"omitempty,min=11,max=15"`
}

type AdminUpdateUserRequest struct {
	Status models.UserStatusEnum `json:"status" validate:"required,oneof=ACTIVE INACTIVE SUSPENDED PENDING"`
	Role   models.RoleEnum       `json:"role" validate:"omitempty,oneof=ADMIN BUSINESS_OWNER RIDER"`
}

type ListUsersRequest struct {
	Role     models.RoleEnum       `form:"role"      validate:"omitempty,oneof=ADMIN BUSINESS_OWNER RIDER"`
	Status   models.UserStatusEnum `form:"status"    validate:"omitempty,oneof=ACTIVE INACTIVE SUSPENDED PENDING"`
	Page     int                   `form:"page"      validate:"omitempty,min=1"`
	PageSize int                   `form:"page_size" validate:"omitempty,min=1,max=100"`
}
