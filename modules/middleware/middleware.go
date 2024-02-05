package middleware

import (
	"github.com/willianrr/hotspot-instance/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeMiddleware() {
	logger = config.GetLogger("middleware")
	db = config.GetPostGresSQL()
}
