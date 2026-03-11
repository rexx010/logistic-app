package models

type BusinessProfile struct {
	Base
	UserID           string         `gorm:"type:uuid;not null;uniqueIndex" json:"user_id"`
	BusinessName     string         `gorm:"not null"                       json:"business_name"`
	BusinessAddress  string         `                                      json:"business_address"`
	SubscriptionPlan string         `gorm:"default:'free'"                 json:"subscription_plan"`
	Status           UserStatusEnum `gorm:"type:varchar(20);default:'ACTIVE'" json:"status"`

	User User `gorm:"foreignKey:UserID"     json:"-"`
	//Deliveries []Delivery `gorm:"foreignKey:BusinessID" json:"deliveries,omitempty"`
}
