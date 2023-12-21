package storage

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Storage struct {
	DB    *gorm.DB
	REDIS redis.Client
}

type IStorage interface {
	InitDb()
	InitCache()
}
