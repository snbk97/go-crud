package server

import (
	"fmt"
	"go_crud/common/constants"
	zlog "go_crud/common/logger"
	"go_crud/common/models"
	"go_crud/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

const (
	FailedToHashPass = "Failed to hash password"
)

func CreateUser(c *gin.Context) {
	var logger = zlog.CreateLogger("CreateUser")
	// facing issue where copilot suggests wrong model to me
	// need to explore how to create static binding of dependencies
	// or maintain enums of models
	userA, _ := models.ModelMetaMap[constants.ModelsE.UserAuthModel].(*models.UserAuthModelMeta)
	userM, _ := models.ModelMetaMap[constants.ModelsE.UserModel].(*models.UserModelMeta)

	var req CreateUserDTO
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		logger.LogErr(err, "CreateUser", "Invalid request body")
		panic("Invalid request body")
	}

	createdUser := userM.CreateUser(&models.User{
		Username: req.Username,
		Email:    req.Email,
		Name:     req.Name,
	})
	utils.HandleOrmError(c, createdUser, "Failed to create user")

	bCryptHashed, bcryptErr := utils.HashPassword(req.Password)
	if bcryptErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": FailedToHashPass,
		})
		return
	}
	createdUserAuth := userA.RegisterUser(
		&models.UserAuth{
			Username: req.Username,
			Password: bCryptHashed,
			Email:    req.Email,
		})
	utils.HandleOrmError(c, createdUserAuth, "Failed to create user auth")

	c.JSON(http.StatusOK, gin.H{
		"message": "User added successfully",
	})
}

func GetSelfUserHandler(c *gin.Context) {
	username := c.MustGet(constants.GinContextE.UserName).(string)
	userM, _ := models.ModelMetaMap[constants.ModelsE.UserModel].(*models.UserModelMeta)

	user := &models.User{}
	record := userM.GetUserByUsername(username, user)
	if record.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found! -> " + username,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"name":     user.Name,
	})

}

func GetUserHandler(c *gin.Context) {
	userM, ok := models.ModelMetaMap[constants.ModelsE.UserModel].(*models.UserModelMeta)
	if !ok {
		panic("userErr::: Failed to get user model")
	}
	username, found := c.Params.Get("username")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	user := &models.User{}
	record := userM.GetUserByUsername(username, user)
	if record.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found! -> " + username,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"name":     user.Name,
	})
}

func DeleteUserHandler(c *gin.Context) {
	userM, _ := models.ModelMetaMap[constants.ModelsE.UserModel].(*models.UserModelMeta)
	userA, _ := models.ModelMetaMap[constants.ModelsE.UserAuthModel].(*models.UserAuthModelMeta)
	username, found := c.Params.Get("username")

	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var user models.User
	var auth models.UserAuth
	record := userM.DeleteUserByUsername(username, user)
	// TODO: figure out how to CASCADE DELETE
	// Maybe use AfterDelete hook?
	authRecord := userA.DeleteUserByUsername(username, auth)
	if record.RowsAffected+authRecord.RowsAffected < 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete user " + username,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User %s deleted successfully", username),
	})
}
