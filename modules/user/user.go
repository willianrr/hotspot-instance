package user

import (
	"github.com/willianrr/hotspot-instance/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeUser() {
	logger = config.GetLogger("user")
	db = config.GetPostGresSQL()
}
