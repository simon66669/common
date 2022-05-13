package predis

import "github.com/go-redis/redis"

func InitRedis() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
