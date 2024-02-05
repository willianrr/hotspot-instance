package user

import (
	"github.com/gin-gonic/gin"
)

func ValidateUser(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	sendSuccess(ctx, "validate-user", user)
}
