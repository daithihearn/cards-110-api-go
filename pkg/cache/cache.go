package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache[T any] struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache[T any](client *redis.Client, ctx context.Context) *RedisCache[T] {
	return &RedisCache[T]{client: client, ctx: ctx}
}

func (c *RedisCache[T]) Set(key string, value T, expiration time.Duration) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(c.ctx, key, jsonData, expiration).Err()
}

func (c *RedisCache[T]) Get(key string) (T, bool, error) {
	var obj T
	val, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return obj, false, err
	}
	if val == "" {
		return obj, false, nil
	}

	err = json.Unmarshal([]byte(val), &obj)
	if err != nil {
		return obj, false, err
	}

	return obj, true, nil
}

func (c *RedisCache[T]) Delete(key string) error {
	return c.client.Del(c.ctx, key).Err()
}
