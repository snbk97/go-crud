package models

import (
	"gorm.io/gorm"
)

type UserModelMeta struct {
	DB *gorm.DB
}

func (meta *UserModelMeta) Init(db *gorm.DB) interface{} {
	return &UserModelMeta{DB: db}
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
}

func (meta *UserModelMeta) CreateUser(u *User) *gorm.DB {
	return meta.DB.Model(&User{}).Create(u)
}

func (meta *UserModelMeta) GetUserByUsername(username string, dest interface{}) *gorm.DB {
	return meta.DB.Model(&User{}).Where("username = ?", username).First(&dest)
}

func (meta *UserModelMeta) DeleteUserByUsername(username string, dest interface{}) *gorm.DB {
	return meta.DB.Model(&User{}).Where("username = ?", username).Where("deleted_at IS NULL").Delete(&dest)
}

// Tried to create a static instance of UserModelMeta, but it doesn't work :(
// var UserModelMetaInstance *UserModelMeta = new(UserModelMeta)
// func (meta *UserModelMeta) Init2(db *gorm.DB) {
// 	meta.DB = db
// }
