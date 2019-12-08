package service

import (
	"bytes"
	"github.com/go-redis/redis/v7"
	"math/rand"
	"time"
)

func (s *service) GetUrl(key string) (url string, err error) {
	url, err = s.dao.GetUrlByKey(key)
	if err == redis.Nil {
		err = ErrorNotExists
	}
	return
}
func (s *service) SaveUrl(keyLen int, url string, expiration time.Duration) (key string, err error) {
	s.Lock()
	defer s.Unlock()
	existsKey, err := s.dao.GetKeyByUrl(url)
	if err != nil {
		if err != redis.Nil {
			return "", err
		}
	}
	if err != redis.Nil {
		return existsKey, nil
	}

	keyBuf := bytes.Buffer{}
	keyCharMap := s.c.Options.KeyCharMap
	keyCharMapLen := len(keyCharMap)
	for i := 0; i < keyLen; i++ {
		index := rand.Intn(keyCharMapLen)
		keyBuf.WriteString(s.c.Options.KeyCharMap[index : index+1])
	}
	key = keyBuf.String()

	if _, err := s.dao.GetUrlByKey(key); err != nil {
		if err != redis.Nil {
			return "", err
		}
		if err != redis.Nil {
			return "", ErrorExists
		}
	}
	err = s.dao.SetUrlByKey(key, url, expiration)
	if err != nil {
		key = ""
	}
	return
}
