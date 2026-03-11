package models

type VehicleTypeEnum string

const (
	Bike     VehicleTypeEnum = "BIKE"
	Car      VehicleTypeEnum = "CAR"
	Tricycle VehicleTypeEnum = "TRICYCLE"
	van      VehicleTypeEnum = "VAN"
)

type RiderProfile struct {
	Base
	UserID          string          `gorm:"type:uuid;not null;uniqueIndex"     json:"user_id"`
	VehicleType     VehicleTypeEnum `gorm:"type:varchar(20)"                   json:"vehicle_type"`
	LicenseNumber   string          `                                           json:"license_number"`
	EarningsBalance float64         `gorm:"type:decimal(12,2);default:0"       json:"earnings_balance"`
	Status          UserStatusEnum  `gorm:"type:varchar(20);default:'PENDING'" json:"status"`
	User            User            `gorm:"foreignKey:UserID"  json:"-"`
	//Deliveries []Delivery `gorm:"foreignKey:RiderID" json:"deliveries,omitempty"`
}
