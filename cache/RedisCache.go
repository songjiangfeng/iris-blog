package cache

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
}

func (c *RedisCache) RedisConnect() (context.Context, *redis.Client) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	return ctx, rdb
}

func (c *RedisCache) RedisGetCache(key string) (val string, err error) {

	ctx, rdb := c.RedisConnect()
	val, err = rdb.Get(ctx, key).Result()
	return val, err
}

func (c *RedisCache) RedisSetCache(key string, data interface{}) (err error) {

	ctx, rdb := c.RedisConnect()
	var json_data []byte

	val, err := c.RedisGetCache(key)

	log.Println(val)

	if err == redis.Nil {
		json_data, err = json.Marshal(data)

		if err != nil {
			panic(err)
		}

		err = rdb.Set(ctx, key, string(json_data), 0).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}
	return err
}
func (c *RedisCache) RedisKeyExists(key string) bool {

	ctx, rdb := c.RedisConnect()
	_, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}
