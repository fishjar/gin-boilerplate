/*
	模型定义
*/

package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Base 给所有模型共用
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"column:id;primary_key;not null"`                   // ID
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;not null"`                // 创建时间
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:update_at;not null"`                 // 更新时间
	DeletedAt *time.Time `json:"-" sql:"index" gorm:"column:deleted_at" binding:"omitempty"` // 软删除时间
}

// BeforeCreate 在创建前给ID赋值一个UUID
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid)
}

// PaginReq 分页查询参数
type PaginReq struct {
	Page uint   `form:"page,default=1"`               // 页码
	Size uint   `form:"size,default=10"`              // 每页数量
	Sort string `form:"sort,default=created_at desc"` // 排序
}

// PaginRes 分页查询结果
type PaginRes struct {
	Page  uint `json:"page" binding:"required"`  // 页码
	Size  uint `json:"size" binding:"required"`  // 每页数量
	Total uint `json:"total" binding:"required"` // 总数
}

// Pagin 分页
func (req PaginReq) Pagin(total uint) PaginRes {
	return PaginRes{
		Page:  req.Page,
		Size:  req.Size,
		Total: total,
	}
}

// BulkDelete 批量删除
type BulkDelete struct {
	IDs []uuid.UUID `form:"ids" json:"ids" binding:"required"`
}

// BulkUpdate 批量更新
type BulkUpdate struct {
	IDs []uuid.UUID            `form:"ids" json:"ids" binding:"required"`
	Obj map[string]interface{} `form:"obj" json:"obj" binding:"required"`
}

// // HTTP 返回信息
// type HTTP struct {
// 	Code    int    `json:"code" binding:"required"`    // 状态码
// 	Message string `json:"message" binding:"required"` // 提示
// }

// HTTPError 错误
type HTTPError struct {
	Code    int      `json:"code" binding:"required"`    // 状态码
	Message string   `json:"message" binding:"required"` // 提示
	Errors  []string `json:"errors"`                     // 详细错误信息
}

// HTTPSuccess 成功
type HTTPSuccess struct {
	Code    int         `json:"code" binding:"required"`    // 状态码
	Message string      `json:"message" binding:"required"` // 提示
	Data    interface{} `json:"data"`                       // 返回数据
}

// HTTPDeleteSuccess 成功
type HTTPDeleteSuccess struct {
	Code    int       `json:"code" binding:"required"`    // 状态码
	Message string    `json:"message" binding:"required"` // 提示
	Data    uuid.UUID `json:"data"`                       // 返回数据
}

// HTTPBulkSuccess 成功
type HTTPBulkSuccess struct {
	Code    int         `json:"code" binding:"required"`    // 状态码
	Message string      `json:"message" binding:"required"` // 提示
	Data    []uuid.UUID `json:"data"`                       // 返回数据
}

// ReqInfo 请求信息
type ReqInfo struct {
	IP     string // IP地址
	Method string // 请求方式
	URI    string // 请求地址
}

// UploadRes 上传文件
type UploadRes struct {
	Origin string `json:"origin" binding:"required"` // 原始文件名
	Type   string `json:"type" binding:"required"`   // 文件类型
	URL    string `json:"url" binding:"required"`    // 文件全路径
	Thumb  string `json:"thumb" binding:"required"`  // 图片缩略图
	Isnew  bool   `json:"isnew" binding:"required"`  // 是否新文件
}

// UploadSuccess 上传单文件
type UploadSuccess struct {
	HTTPSuccess
	Data UploadRes `json:"data" binding:"required"`
}

// UploadMultiSuccess 上传多文件
type UploadMultiSuccess struct {
	HTTPSuccess
	Data []UploadRes `json:"data" binding:"required"`
}

// BeforeUpdate 钩子
func (base *Base) BeforeUpdate() (err error) {
	fmt.Println("-------BeforeUpdate Hooks--------")
	fmt.Println(base.ID)
	return
}

// BeforeDelete 钩子
func (base *Base) BeforeDelete() (err error) {
	fmt.Println("-------BeforeDelete Hooks--------")
	fmt.Println(base.ID)
	return
}
