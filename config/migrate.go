package config

import (
	"log"
	"logisticApp/data/models"
)

func MigrateDB() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.BusinessProfile{},
		&models.RiderProfile{},
		&models.Delivery{},
		&models.Payment{},
		&models.Invoice{},
		&models.Commission{},
		&models.Notification{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully")
}
