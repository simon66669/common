package predis

import "github.com/go-redis/redis"

func InitRedis() (*redis.Client, error) {

	Rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := Rdb.Ping().Result()

	if err != nil {
		return nil, err
	}

	return Rdb, nil
}
