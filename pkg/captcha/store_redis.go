package captcha

import (
	"errors"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}

	return nil
}

func (s *RedisStore) Get(key string, clear bool) string {
	val := s.RedisClient.Get(s.KeyPrefix + key)

	if clear {
		s.RedisClient.Del(s.KeyPrefix + key)
	}

	return val
}

func (s RedisStore) Verify(key string, answer string, clear bool) bool {
	val := s.Get(key, clear)
	return val == answer
}
