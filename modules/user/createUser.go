package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willianrr/hotspot-instance/schemas"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx *gin.Context) {
	request := LoginRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 4)
	if err != nil {
		logger.ErrorF("Hash not validate error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Username:   "",
		Email:      request.Email,
		Phone:      "",
		FirstLogin: true,
		Firstname:  "",
		Lastname:   "",
		Role:       "",
		Password:   string(hasPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		logger.ErrorF("Error User: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error create user please try again")
		return
	}

	sendSuccess(ctx, "create-user", user)
}
