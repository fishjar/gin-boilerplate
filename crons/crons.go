package crons

import "github.com/robfig/cron"

// 定时任务
const (
	EVERY10SENCOND = "@every 10s" // 每10秒
	EVERY20SENCOND = "@every 20s" // 每20秒
	EVERY30SENCOND = "@every 30s" // 每30秒
)

// Cron 定时任务
var Cron *cron.Cron

func init() {
	c := cron.New()
	c.AddJob(EVERY10SENCOND, &TestJob{"10s"})
	c.AddFunc(EVERY20SENCOND, TestFunc("20s"))
	c.Start()
	Cron = c
}
