package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

const apqPrefix = "apq:"

func New(redisAddress string, password string, ttl time.Duration) (*cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddress,
		DB:   0,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("could not ping redis: %w", err)
	}

	return &cache{client: client, ttl: ttl}, nil
}

func (c *cache) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(apqPrefix+key, value, c.ttl)
}

func (c *cache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.client.Get(apqPrefix + key).Result()
	if err != nil {
		return "", false
	}
	return s, true
}
