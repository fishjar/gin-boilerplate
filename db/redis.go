package db

import (
	"fmt"

	"github.com/fishjar/gin-boilerplate/config"
	"github.com/go-redis/redis/v7"
)

// Redis 全局实例
var Redis *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,     // redis 地址
		Password: config.Config.Redis.Password, // redis 密码
		DB:       config.Config.Redis.Name,     // use default DB
	})

	if pong, err := client.Ping().Result(); err != nil {
		fmt.Println("----redis ping----", pong, err)
		panic("redis连接错误")
	}

	Redis = client
}
