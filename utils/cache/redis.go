package cache

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type CacheRedis struct {
	Conn   redis.Conn
	Prefix string
}

func (r *CacheRedis) Set(key string, val string) bool {
	conn := r.Conn
	defer conn.Close()

	conn.Do("SET", r.getKey(key), val)
	return true

}

func (r *CacheRedis) Get(key string) (string, error) {
	conn := r.Conn
	defer conn.Close()

	s, err := redis.String(conn.Do("GET", r.getKey(key)))
	if err != nil {
		fmt.Println(err)
		return "", errors.New("redis error:" + err.Error())
	}

	return s, nil
}

func (r *CacheRedis) getKey(key string) string {

	if r.Prefix == "" {
		return key
	} else {
		key = r.Prefix + "_" + key
		return key
	}
}
