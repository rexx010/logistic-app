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
	Name           string         `gorm:"not null" json:"name"`
	Email          string         `gorm:"uniqueIndex;not null" json:"email"`
	Phone          string         `gorm:"uniqueIndex" json:"phone"`
	Password       string         `gorm:"not null" json:"-"`
	ProfilePicture string         `gorm:"default:null" json:"profile_picture"`
	Role           RoleEnum       `gorm:"type:varchar(20);not null" json:"role"`
	Status         UserStatusEnum `gorm:"type:varchar(20);default:'ACTIVE'" json:"status"`

	BusinessProfile *BusinessProfile `gorm:"foreignKey:UserID" json:"businessProfile,omitempty"`
	RiderProfile    *RiderProfile    `gorm:"foreignKey:UserID" json:"riderProfile,omitempty"`
	Notifications   []Notification   `gorm:"foreignKey:UserID" json:"notifications,omitempty"`
}
