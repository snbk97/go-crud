package models

import (
	"go_crud/common/utils"

	"gorm.io/gorm"
)

type UserAuthModelMeta struct {
	DB *gorm.DB
}

func (meta *UserAuthModelMeta) Init(db *gorm.DB) interface{} {
	return &UserAuthModelMeta{DB: db}
}

type UserAuth struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;primaryKey"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"not null;unique"`
}

func (ua *UserAuth) TableName() string {
	return "user_auth"
}

func (meta *UserAuthModelMeta) RegisterUser(u *UserAuth) *gorm.DB {
	return meta.DB.Model(&UserAuth{}).Create(u)
}

func (meta *UserAuthModelMeta) DeleteUserByUsername(username string, dest interface{}) *gorm.DB {
	return meta.DB.Model(&UserAuth{}).Where("username = ?", username).Where("deleted_at IS NULL").Delete(&dest)
}

func (meta *UserAuthModelMeta) Login(username string, password string, dest interface{}) (record *gorm.DB, loginSuccess bool) {
	loginSuccess = false
	userAuth := &UserAuth{}

	record = meta.DB.Model(&UserAuth{}).Where("username = ? AND deleted_at IS NULL", username).First(userAuth)
	if record.Error != nil {
		return record, loginSuccess
	}

	loginSuccess = utils.VerifyPassword(userAuth.Password, password)

	return
}
