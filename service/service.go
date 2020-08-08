/*
	服务封装
*/

package service

import (
	"net/http"

	"github.com/fishjar/gin-boilerplate/logger"
	"github.com/fishjar/gin-boilerplate/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewError 创建错误对象
func NewError(msg string, code int, errs []error, c *gin.Context) model.HTTPError {
	var errStrs []string
	for _, err := range errs {
		errStrs = append(errStrs, err.Error())
	}
	if c == nil {
		go logger.Log.WithFields(logrus.Fields{
			"code": code,
			"errs": errs,
		}).Warn(msg)
	} else {
		go logger.Log.WithFields(logrus.Fields{
			"code": code,
			"errs": errs,
			"req":  GetReqInfo(c),
		}).Warn(msg)
	}
	return model.HTTPError{
		Code:    code,
		Message: msg,
		Errors:  errStrs,
	}

}

// GetReqInfo 获取请求信息
func GetReqInfo(c *gin.Context) model.ReqInfo {
	reqMethod := c.Request.Method  // 请求方式
	reqURI := c.Request.RequestURI // 请求路由
	clientIP := c.ClientIP()       // 请求IP
	return model.ReqInfo{
		Method: reqMethod,
		URI:    reqURI,
		IP:     clientIP,
	}
}

// HTTPError 返回错误
func HTTPError(c *gin.Context, msg string, code int, err error) {
	var httpError model.HTTPError
	if err == nil {
		httpError = NewError(msg, code, []error{}, c)
	} else {
		httpError = NewError(msg, code, []error{err}, c)
	}
	c.JSON(code, httpError)
}

// HTTPAbortError 返回错误
func HTTPAbortError(c *gin.Context, msg string, code int, err error) {
	var httpError model.HTTPError
	if err == nil {
		httpError = NewError(msg, code, []error{}, c)
	} else {
		httpError = NewError(msg, code, []error{err}, c)
	}
	c.AbortWithStatusJSON(code, httpError)
}

// HTTPSuccess 返回成功
func HTTPSuccess(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, model.HTTPSuccess{
		Message: msg,
	})
}
