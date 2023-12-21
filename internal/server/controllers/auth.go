package server

import (
	"go_crud/common/constants"
	zlog "go_crud/common/logger"
	models "go_crud/common/models"
	storage "go_crud/common/storage"
	"go_crud/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginDTO struct {
	Username string
	Password string
}

type LogoutDTO struct {
	Username string
}

const (
	InvalidRequestBodyError      = "Invalid request body"
	InvalidUsernamePasswordError = "Invalid username or password"
	ErrorGeneratingToken         = "Error generating token"
)

var logger = zlog.CreateLogger("AuthLoginHandler")

func AuthLoginHandler(c *gin.Context) {

	var req LoginDTO
	if utils.CreateHttpError(c, c.BindJSON(&req), http.StatusUnauthorized, "Invalid request body") {
		return
	}

	userA, _ := models.ModelMetaMap[constants.ModelsE.UserAuthModel].(*models.UserAuthModelMeta)
	record, loginSuccess := userA.Login(req.Username, req.Password, &req)
	if utils.CreateHttpError(c, record.Error, http.StatusUnauthorized, InvalidUsernamePasswordError) || record.RowsAffected == 0 {
		return
	}

	if !loginSuccess {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": InvalidUsernamePasswordError,
		})
		return
	}

	token, err := utils.GetAuthJwt(req.Username)
	if utils.CreateHttpError(c, err, http.StatusUnauthorized, ErrorGeneratingToken) {
		return
	}

	storage.REDIS.Set(c, req.Username, token, constants.AUTH_TTL)
	logger.LogInfo("Login successful -> ", req.Username, "1234")
	c.SetCookie("token", token, int(constants.AUTH_TTL.Microseconds()), "/", "", false, true)
	c.JSON(200, gin.H{
		"message": "Login successful -> " + req.Username,
	})
}

func AuthLogoutHandler(c *gin.Context) {
	var req LogoutDTO
	if utils.CreateHttpError(c, c.Bind(&req), http.StatusBadRequest, InvalidRequestBodyError) {
		return
	}
	storage.REDIS.Del(c, req.Username)
	c.SetCookie("token", "", 0, "/", "", false, true)
	c.JSON(200, gin.H{
		"message": "Logout successful -> " + req.Username,
	})
}
