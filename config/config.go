package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
	radius *gorm.DB
)

func Init() error {
	var err error

	// Initialize PostgreSQL
	db, err = InitializePostgreSQl()

	if err != nil {
		return fmt.Errorf("Error initialize PostgreSql %v", err)
	}
	return nil
}

func GetPostGresSQL() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}
