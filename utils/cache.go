package utils

import (
	"time"
	"wegou/config"
	"wegou/utils/cache"

	"github.com/gomodule/redigo/redis"
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

var (
	pool *redis.Pool
)

func Redis(web string) (r *cache.Redis) {
	redisConfig := config.GetRedisConfig(web)
	pool = newPool(redisConfig.Server, redisConfig.Auth, redisConfig.Db)
	r = &cache.Redis{Conn: pool.Get(), Prefix: web}
	return r
}
