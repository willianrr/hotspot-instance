package user

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/willianrr/hotspot-instance/schemas"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(ctx *gin.Context) {
	request := LoginRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}

	if err := db.First(&user, "email = ?", request.Email).Error; err != nil {
		logger.ErrorF("Invalid email or password: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Invalid email or password")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		logger.ErrorF("Invalid email or password: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":         user.ID,
		"exp":         time.Now().Add(time.Hour * 24 * 30).Unix(),
		"first_login": user.FirstLogin,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		logger.ErrorF("Fail to create token: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	sendSuccess(ctx, "login-user", true)
}
