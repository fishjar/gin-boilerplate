package tasks

import "github.com/hibiken/asynq"

// NewEmailDeliveryTask 创建发送邮件任务
func NewEmailDeliveryTask(userID int, tmplID string) *asynq.Task {
	payload := map[string]interface{}{"user_id": userID, "template_id": tmplID}
	return asynq.NewTask(emailDelivery, payload)
}

// NewImageProcessingTask 创建图片处理任务
func NewImageProcessingTask(src, dst string) *asynq.Task {
	payload := map[string]interface{}{"src": src, "dst": dst}
	return asynq.NewTask(imageProcessing, payload)
}
