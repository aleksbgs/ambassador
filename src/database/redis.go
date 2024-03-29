package database

import (
	"context"
	"fmt"
	"github.com/aleksbgs/ambassador/utils"
	"github.com/go-redis/redis/v8"
	"time"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: utils.ViperEnvVariable("REDIS") + ":6379",
		DB:   0,
	})
}

func SetupCacheChannel() {
	CacheChannel = make(chan string)

	go func(ch chan string) {
		for {
			time.Sleep(5 * time.Second)

			key := <-ch

			Cache.Del(context.Background(), key)

			fmt.Println("Cache cleared " + key)
		}
	}(CacheChannel)
}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key
	}
}
