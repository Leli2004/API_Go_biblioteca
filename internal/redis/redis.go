package redis

import (
	"fmt"

	"github.com/Leli2004/API_Go_biblioteca/config"
	"github.com/redis/go-redis/v9"
)

func NewClient(cfg config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.Database,
	})
}
