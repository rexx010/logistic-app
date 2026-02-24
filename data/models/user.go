package models

type RoleEnum string

const (
	Admin         RoleEnum = "ADMIN"
	BusinessOwner RoleEnum = "BUSINESS_OWNER"
	Rider         RoleEnum = "RIDER"
)

type UserStatusEnum string

const (
	Active    UserStatusEnum = "ACTIVE"
	Inactive  UserStatusEnum = "INACTIVE"
	Suspended UserStatusEnum = "SUSPENDED"
	Pending   UserStatusEnum = "PENDING"
)

type User struct {
	Base
	Name     string         `gormm:"not null" json:"name"`
	Email    string         `gormm:"uuniqueIndex;not null" json:"email"`
	Phone    string         `gormm:"uuniqueIndex" json:"phone"`
	Password string         `gormm:"not null" json:"-"`
	Role     RoleEnum       `gorm:"type:varchar(20);not null" json:"role"`
	Status   UserStatusEnum `gorm:"type:varchar(20);default:'ACTIVE'" json:"status"`

	BusinessProfile *BusinessProfile `gorm:"foreignKey:UserID" json:"businessProfile,omitempty"`
	RiderProfile    *RiderProfile    `gorm:"foreignKey:UserID" json:"riderProfile,omitempty"`
	Notifications   []Notification   `gorm:"foreignKey:UserID" json:"notifications,omitempty"`
}
