package database

import (
	"log"

	"github.com/shohan-joarder/go_pos/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Connected to PostgreSQL")

	// Auto-migrate the models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&models.Role{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&models.RolePermission{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&models.Warehouse{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
