package cache

import (
	"blog/conf"
	uredis "blog/util/redis"

	"github.com/go-redis/redis/v7"
)

//Cache 层
type Cache struct {
	c     *conf.Config
	Redis *redis.Client
}

//NewCache 实例化 Dao
func NewCache(c *conf.Config) *Cache {
	return &Cache{
		c:     c,
		Redis: uredis.NewRedis(c.Redis),
	}
}
