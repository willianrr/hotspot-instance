package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/willianrr/hotspot-instance/config"
	"github.com/willianrr/hotspot-instance/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	fmt.Println(gin.Version)
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
