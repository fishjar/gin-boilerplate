package tasks

import (
	"github.com/fishjar/gin-boilerplate/config"
	"github.com/hibiken/asynq"
)

// Client 全局客户端
var Client *asynq.Client

func init() {
	r := asynq.RedisClientOpt{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.Name,
	}
	Client = asynq.NewClient(r)
}
