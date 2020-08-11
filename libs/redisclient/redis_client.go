package redisclient

import (
	"time"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/configure"
)

var (
	client *redis.Pool
)

func GetConn() redis.Conn {
	if client == nil {
		log.Info("Redis init")
		config := configure.New().RedisServer
		client = &redis.Pool{
			MaxIdle:     config.MaxIdle,
			MaxActive:   config.MaxActive,
			IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", config.Host+config.Port)
				if err != nil {
					return nil, err
				}
				c.Do("SELECT", config.DB)
				return c, nil
			},
		}
		log.Debugln("Redis init success", config.Host+config.Port)

	}
	return client.Get()
}
