package utils

import (
	"encoding/json"
	"errors"
	"time"
	"wegou/config"
	"wegou/types"
	"wegou/utils/cache"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

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

func GetCache(web string) (r cache.Cache) {

	toolsConfig := config.GetToolsConfig()
	switch toolsConfig.Cache {
	case "redis":
		redisConfig := config.GetRedisConfig()
		pool = newPool(redisConfig.Server, redisConfig.Auth, redisConfig.Db)
		r = &cache.CacheRedis{Conn: pool.Get(), Prefix: web}
	}

	return r
}

//获取公众号缓存
func GetWechatConfig(web string) (wechat types.Db) {
	jsonAccount, err := GetCache(web).Get("wechat")
	if err != nil {
		logrus.Error(errors.New("json account error:" + err.Error()))
		return wechat
	}
	if jsonAccount != "" {
		json.Unmarshal([]byte(jsonAccount), &wechat)
	}
	return wechat
}
