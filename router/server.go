package router

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fishjar/gin-boilerplate/config"
	"github.com/fishjar/gin-boilerplate/crons"
	"github.com/fishjar/gin-boilerplate/db"
	"github.com/fishjar/gin-boilerplate/logger"
	"github.com/fishjar/gin-boilerplate/tasks"
)

// RunGinServer 运行gin服务
func RunGinServer(taskDone, allDone chan bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("gin revocer")
		}
	}()

	defer db.DB.Close()              // 关闭数据库连接
	defer db.Redis.Close()           // 关闭Redis连接
	defer logger.LogFile.Close()     // 关闭日志文件
	defer logger.LogGinFile.Close()  // 关闭日志文件
	defer logger.LogReqFile.Close()  // 关闭日志文件
	defer logger.LogGormFile.Close() // 关闭日志文件
	defer crons.Cron.Stop()          // 关闭定时任务

	r := InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Config.HTTPPort),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() { // gin关闭线程
		<-taskDone // 阻塞，等待退出信号
		now := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil { // 退出gin服务
			fmt.Println("关闭gin服务失败", err)
		}
		select {
		case <-ctx.Done(): // 阻塞，等待3秒
			fmt.Println("----timeout of 3 seconds-----")
		}
		fmt.Println("------gin exited--------", time.Since(now))
		allDone <- true
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed { //阻塞，等待关闭或错误
		fmt.Println(err)
		panic("gin启动失败")
	}
}

// RunTaskServer 运行task服务
func RunTaskServer(taskDone chan bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("tasks sever revocer")
		}
	}()

	defer tasks.Client.Close() // 关闭任务队列服务

	if err := tasks.Server(); err != nil { // 阻塞，等待退出信号
		fmt.Println(err)
		panic("任务队列启动失败")
	}
	fmt.Println("------tasks exited--------")
	taskDone <- true
}
