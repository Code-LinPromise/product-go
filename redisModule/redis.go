package redisModule

import "github.com/redis/go-redis/v9"

var Redis *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.21.80.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Redis = client
}
