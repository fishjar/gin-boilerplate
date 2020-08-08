package crons

import (
	"fmt"
)

// TestJob 测试定时任务
type TestJob struct {
	Value string
}

// Run 实现定时任务接口
func (job *TestJob) Run() {
	fmt.Println("测试定时任务:job", job.Value)
}

// TestFunc 测试定时函数
func TestFunc(s string) func() {
	return func() {
		fmt.Println("测试定时任务:func", s)
	}
}
