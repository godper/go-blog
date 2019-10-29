package cache

import (
	"encoding/json"
	"fmt"
	"time"
)

//Set set
func (c *Cache) Set(key string, value interface{}, expire time.Duration) {
	value, err := json.Marshal(value)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	err = c.Redis.Set(key, value, expire).Err()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}

//Get set
func (c *Cache) Get(key string) ([]byte, error) {
	val, err := c.Redis.Get(key).Result()
	return []byte(val), err
}

//DeleteLike 删除键值对like
func (c *Cache) DeleteLike(key string) error {
	keys, err := c.Redis.Keys(key).Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		if err := c.Redis.Del(key).Err(); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

//Delete 删除键
func (c *Cache) Delete(key string) error {
	if err := c.Redis.Del(key).Err(); err != nil {
		return err
	}
	return nil
}
