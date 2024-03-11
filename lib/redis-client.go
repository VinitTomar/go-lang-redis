package lib

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func GetRedisClient() *redis.Client {
	if client == nil {
		opt, err := redis.ParseURL(os.Getenv("REDIS_CONNECT_URL"))

		if err != nil {
			panic(err)
		}

		client = redis.NewClient(opt)
	}

	return client
}
