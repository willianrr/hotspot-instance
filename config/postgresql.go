package config

import (
	"github.com/willianrr/hotspot-instance/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQl() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.ErrorF("postgresSQL opening error: %v", err)
		return nil, err
	}
	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		logger.ErrorF("postgresSQL automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
