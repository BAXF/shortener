package models

import (
	"context"
	redis "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func ConnectRedis(addr string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisClient{Client: client}
}

func (r *RedisClient) GetVisits(ctx context.Context, shortURL string) (int, error) {
	visits, err := r.Client.Get(ctx, shortURL).Int()
	if err != nil {
		return 0, err
	}
	return visits, nil
}
