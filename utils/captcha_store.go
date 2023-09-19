package utils

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"time"
)

var captchaRedisExpiration = 15 * time.Minute

type captchaRedisStore struct {
	client     *redis.Client // redis连接
	expiration time.Duration // 有效时长
	keyPrefix  string        // 存储键的前缀
}

func NewCaptchaRedisStore(client *redis.Client, keyPrefix string) base64Captcha.Store {
	return &captchaRedisStore{
		client:     client,
		expiration: captchaRedisExpiration,
		keyPrefix:  keyPrefix,
	}
}

func (c *captchaRedisStore) SetExpiration(expiration time.Duration) {
	c.expiration = expiration
}

func (c *captchaRedisStore) Set(id string, value string) error {
	err := c.client.Set(context.Background(), c.keyPrefix+id, value, c.expiration*time.Second).Err()
	return err
}

func (c *captchaRedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	var value string
	var err error
	if clear {
		value, err = c.client.GetDel(ctx, c.keyPrefix+id).Result()
	} else {
		value, err = c.client.Get(ctx, c.keyPrefix+id).Result()
	}
	if err != nil {
		return ""
	}
	return value
}

func (c *captchaRedisStore) Verify(id, answer string, clear bool) (match bool) {
	return c.Get(id, clear) == answer
}
