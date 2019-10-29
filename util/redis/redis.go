package uredis

import (
	"time"
	"github.com/go-redis/redis/v7"
)

//RedisConf Redis 配置
type RedisConf struct {
	DB          int
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

//Redis缓存客户端单例

//NewRedis New redis
func NewRedis(conf *RedisConf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: conf.Password,
		DB:       conf.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
