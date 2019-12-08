package conf

import "github.com/go-redis/redis/v7"

type Config struct {
	DB struct {
		Redis *redis.Options
	}

	Options struct {
		HttpAddr    []string
		RedisPrefix string
		Security    struct {
			Token []string
		}
		KeyCharMap string
		KeyLength  int
		RetryCount int
	}
}
