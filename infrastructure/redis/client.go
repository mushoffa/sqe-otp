package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService interface {
	Set(context.Context, string, any, time.Duration) error
	Exists(context.Context, string) (bool, error)
	Del(context.Context, string) error
	Close() error
}

type client struct {
	instance *redis.Client
}

func New(address string) RedisService {
	conn := redis.NewClient(&redis.Options{
		Addr: address,
	})

	db := &client{
		instance: conn,
	}

	return db
}

func (c *client) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return c.instance.Set(ctx, key, value, expiration).Err()
}

func (c *client) Exists(ctx context.Context, key string) (bool, error) {
	exists, err := c.instance.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if exists > 0 {
		return true, nil
	}

	return false, nil
}

func (c *client) Del(ctx context.Context, key string) error {
	return c.instance.Del(ctx, key).Err()
}

func (c *client) Close() error {
	return c.instance.Close()
}
