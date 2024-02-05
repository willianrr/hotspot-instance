package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/willianrr/hotspot-instance/schemas"
)

func RequireAuth(ctx *gin.Context) {

	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		sendError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			logger.ErrorF("expirated: %v", err.Error())
			sendError(ctx, http.StatusUnauthorized, err.Error())
			return
		}
		user := schemas.User{}

		if err := db.First(&user, claims["sub"]).Error; err != nil {
			logger.ErrorF("Not found user: %v", err.Error())
			sendError(ctx, http.StatusUnauthorized, "Not found user")
			return
		}

		ctx.Set("user", user.ID)

		ctx.Next()

	} else {
		sendError(ctx, http.StatusUnauthorized, err.Error())
	}

}
