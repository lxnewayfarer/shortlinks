package storage

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	TxPipeline() redis.Pipeliner
}

func Nil() any {
	return redis.Nil
}

func InitRedis() (*redis.Client, error) {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return nil, err
	}

	return redis.NewClient(opts), nil
}

func InitMockRedis() (*redis.Client, redismock.ClientMock) {
	db, mock := redismock.NewClientMock()
	return db, mock
}