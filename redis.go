package vps_monitoring_notifier

import (
	"fmt"

	"github.com/go-redis/redis"
)

func pingRedis(redisAddr string) error {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return fmt.Errorf("ping: %w", err)
	}

	err = client.Close()
	if err != nil {
		return fmt.Errorf("close: %w", err)
	}

	return nil
}
