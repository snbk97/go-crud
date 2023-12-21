package storage

import (
	"github.com/redis/go-redis/v9"
)

func InitCache() {
	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if cache == nil {
		panic("failed to connect cache")
	}
	REDIS = *cache
}
