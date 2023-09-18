package utils

import (
	"context"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"time"
)

var CaptchaRedisExpiration = 15 * time.Minute

type captchaRedisStore struct {
	client     *redis.Client
	expiration time.Duration // 有效时长
}

func NewCaptchaRedisStore(client *redis.Client) base64Captcha.Store {
	return &captchaRedisStore{
		client:     client,
		expiration: CaptchaRedisExpiration,
	}
}

func (c *captchaRedisStore) SetExpiration(expiration time.Duration) {
	c.expiration = expiration
}

func (c *captchaRedisStore) Set(id string, value string) error {
	err := c.client.Set(context.Background(), id, value, c.expiration*time.Second).Err()
	return err
}

func (c *captchaRedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	var value string
	var err error
	if clear {
		value, err = c.client.GetDel(ctx, id).Result()
	} else {
		value, err = c.client.Get(ctx, id).Result()
	}
	if err != nil {
		return ""
	}
	return value
}

func (c *captchaRedisStore) Verify(id, answer string, clear bool) (match bool) {
	return c.Get(id, clear) == answer
}
