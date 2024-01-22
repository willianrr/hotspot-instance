package handler

import (
	"github.com/willianrr/hotspot-go/config"

	"gorm.io/gorm"
)

var (
	logger 	*config.Logger
	db		*gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetPostGresSQL()
}