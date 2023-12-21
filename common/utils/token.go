package utils

import (
	"errors"
	"go_crud/common/constants"
	zlog "go_crud/common/logger"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var logger = zlog.CreateLogger("tokenUtils")

func GetAuthJwt(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "go-crud"
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Duration(constants.AUTH_TTL)).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		errMsg := "failed to generate JWT token"
		logger.LogErr(err, errMsg)
		return "", errors.New(errMsg)
	}
	return tokenString, nil
}
