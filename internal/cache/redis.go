package cache

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/mohammad-hosein-shahpouri/uptime-phoenix/internal/config"
)

type Cache struct {
	client *redis.Client
}

func New(config *config.RedisConfig) (*Cache, error) {
	redisURL := fmt.Sprintf("redis://%s:%s@%s:%s/%d", config.User, config.Password, config.Host, config.Port, config.Database)

	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("invalid redis url: %w", err)
	}

	client := redis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	slog.Info("Connected to Redis successfully", "addr", opts.Addr)
	return &Cache{client: client}, nil
}

func (c *Cache) Close() error {
	return c.client.Close()
}

func (c *Cache) Client() *redis.Client {
	return c.client
}
