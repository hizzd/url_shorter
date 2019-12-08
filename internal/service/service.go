package service

import (
	"errors"
	"sync"
	"time"
	"url_short/internal/conf"
	"url_short/internal/dao"
)

var (
	ErrorExists    = errors.New("item exists")
	ErrorNotExists = errors.New("item not exists")
)

type Service interface {
	GetUrl(key string) (url string, err error)
	SaveUrl(keyLen int, url string, expiration time.Duration) (key string, err error)
}

func NewService(c *conf.Config) Service {
	return &service{
		c:   c,
		dao: dao.NewDao(c),
	}
}

type service struct {
	sync.Mutex
	c   *conf.Config
	dao dao.Dao
}
