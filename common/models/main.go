package models

import "gorm.io/gorm"

var ModelMetaMap = make(map[string]interface{})

type ModelMeta interface {
	Init(db *gorm.DB) interface{}
}
