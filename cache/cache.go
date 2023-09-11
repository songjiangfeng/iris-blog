package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
}

func (c *RedisCache) redisConnect() (context.Context, *redis.Client) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	return ctx, rdb
}

func (c *RedisCache) redisCache(key string, data interface{}, model interface{}) (obj struct{}, err error) {

	ctx, rdb := c.redisConnect()

	json_data, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		err := rdb.Set(ctx, key, json_data, 0).Err()
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(val), &obj)
	if err != nil {
		panic(err)
	}
	return obj, err
}
