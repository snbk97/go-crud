package server

import (
	"errors"
	"go_crud/common/constants"
	zlog "go_crud/common/logger"
	storage "go_crud/common/storage"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var logger = zlog.CreateLogger("ValidateUserMiddleware")

const (
	InvalidTokenError = "Invalid token"
	ParseTokenError   = "Failed to parse token"
	FetchTokenError   = "Failed to fetch token"
	InvalidClaims     = "Invalid claims"
	TokenExpiredError = "Token expired"
	JWTSecretEnv      = "JWT_SECRET"
)

func ValidateUserMiddleware(c *gin.Context) {
	// get token from cookie
	cookieToken, cookieErr := c.Request.Cookie("token")
	if cookieErr != nil {
		handleMiddlewareError(c, cookieErr, FetchTokenError)
		return
	}

	claims, clErr := parseTokenAndVerify(c, cookieToken)
	if clErr != nil {
		handleMiddlewareError(c, clErr, InvalidTokenError)
		return
	}

	user := claims["username"].(string)
	if user == "" {
		handleMiddlewareError(c, nil, InvalidClaims)
		return
	}

	// checks validity in redis cache
	result := storage.REDIS.Get(c, user)
	if len(result.Val()) == 0 {
		handleMiddlewareError(c, nil, TokenExpiredError)
		return
	}

	// after this point, we know the user is valid
	// set the username in the gin context
	c.Set(constants.GinContextE.UserName, user)

	c.Next()
}

func parseTokenAndVerify(c *gin.Context, cookieToken *http.Cookie) (jwt.MapClaims, error) {
	token, err := jwt.Parse(cookieToken.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(JWTSecretEnv)), nil
	})
	if err != nil {
		return nil, errors.New(strings.ToLower(ParseTokenError))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New(strings.ToLower(InvalidClaims))
	}

	return claims, nil
}

func handleMiddlewareError(c *gin.Context, err error, message string) {
	logger.LogErr(err, message)
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": message,
	})
	c.Abort()

}
