package redis

import (
	"Library_Project/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

var (
	Client *redis.Client
)

func Init(conf *config.RedisConfig) (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password:     conf.Password,
		DB:           conf.DB,
		PoolSize:     conf.PollSize,
		MinIdleConns: conf.MinIdleConns,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err = Client.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("redis请求ping失败了，请重新检查一下")
		return
	}
	return nil
}

func Close() {
	_ = Client.Close()
}
