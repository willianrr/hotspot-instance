package handler

import (
	"github.com/willianrr/hotspot-instance/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetPostGresSQL()
}
