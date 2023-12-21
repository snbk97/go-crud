package storage

import (
	"go_crud/common/constants"
	models "go_crud/common/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() {
	_db, err := gorm.Open(
		sqlite.Open("go_crud.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	DB = _db
}

func RegisterModels() {
	modelList := []interface{}{
		&models.User{},
		&models.UserAuth{},
		&models.Post{},
	}
	DB.AutoMigrate(modelList...)
}

/**
 * Injects the depencies into the models
 * Stores the models in memory
 */
func ConnectModels() {
	models.ModelMetaMap[constants.ModelsE.UserModel] = (&models.UserModelMeta{}).Init(DB)
	models.ModelMetaMap[constants.ModelsE.UserAuthModel] = (&models.UserAuthModelMeta{}).Init(DB)
	models.ModelMetaMap[constants.ModelsE.PostModel] = (&models.PostModelMeta{}).Init(DB)
}
