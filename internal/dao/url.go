package dao

import (
	"fmt"
	"time"
)

const (
	_keyUrlByKey = "%skey:%s"
	_keyKeyByUrl = "%surl:%s"
)

func (d *dao) GetUrlByKey(key string) (url string, err error) {
	return d.rc.Get(fmt.Sprintf(_keyUrlByKey, d.c.Options.RedisPrefix, key)).Result()
}
func (d *dao) GetKeyByUrl(url string) (key string, err error) {
	return d.rc.Get(fmt.Sprintf(_keyKeyByUrl, d.c.Options.RedisPrefix, url)).Result()

}

func (d *dao) SetUrlByKey(key string, url string, expiration time.Duration) error {
	err := d.rc.Set(fmt.Sprintf(_keyUrlByKey, d.c.Options.RedisPrefix, key), url, expiration).Err()
	if err != nil {
		return err
	}
	return d.rc.Set(fmt.Sprintf(_keyKeyByUrl, d.c.Options.RedisPrefix, url), key, expiration).Err()
}
