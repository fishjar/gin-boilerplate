package tasks

import (
	"log"

	"github.com/fishjar/gin-boilerplate/config"
	"github.com/hibiken/asynq"
)

// Server 启动服务
func Server() error {
	r := asynq.RedisClientOpt{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.Name,
	}
	srv := asynq.NewServer(r, asynq.Config{
		// Specify how many concurrent workers to use
		Concurrency: 10,
		// Optionally specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
		// See the godoc for other configuration options
	})

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(emailDelivery, handleEmailDeliveryTask)
	mux.Handle(imageProcessing, newImageProcessor())
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
		return err
	}

	return nil
}
