// @Author	zhangjiaozhu 2023/3/18 16:00:00
package redis

import "github.com/go-redis/redis/v8"

var RDB *redis.Client

func InitRedisDb() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return nil
}
