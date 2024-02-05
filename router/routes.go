package router

import (
	"github.com/willianrr/hotspot-instance/modules/handler"
	"github.com/willianrr/hotspot-instance/modules/middleware"
	"github.com/willianrr/hotspot-instance/modules/user"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	handler.InitializeHandler()
	user.InitializeUser()
	middleware.InitializeMiddleware()
	{
		// Openings
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)

	}
	{
		// Users
		v1.POST("/user", user.CreateUser)
		v1.POST("/login", user.LoginUser)
		v1.GET("/validate", middleware.RequireAuth, user.ValidateUser)
	}
}
