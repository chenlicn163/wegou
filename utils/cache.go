package utils

import (
	"time"
	"wegou/config"
	"wegou/utils/cache"

	"github.com/gomodule/redigo/redis"
)

//缓存
type Cache interface {
	Set(key string, val string) bool
	Get(key string) (string, error)
}

var (
	pool *redis.Pool
)

func newPool(server string, password string, db int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

//获取缓存
func GetCache(web string) (r Cache) {

	toolsConfig := config.GetToolsConfig()
	switch toolsConfig.Cache {
	case "redis":
		redisConfig := config.GetRedisConfig()
		pool = newPool(redisConfig.Server, redisConfig.Auth, redisConfig.Db)
		r = &cache.CacheRedis{Conn: pool.Get(), Prefix: web}
	}

	return r
}
