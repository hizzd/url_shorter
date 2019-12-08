package dao

import (
	"github.com/go-redis/redis/v7"
	"time"
	"url_short/internal/conf"
)

type Dao interface {
	GetUrlByKey(key string) (url string, err error)
	GetKeyByUrl(url string) (key string, err error)
	SetUrlByKey(key string, url string, expiration time.Duration) error
}

func NewDao(c *conf.Config) Dao {
	return &dao{
		c:  c,
		rc: redis.NewClient(c.DB.Redis),
	}
}

type dao struct {
	c  *conf.Config
	rc *redis.Client
}
