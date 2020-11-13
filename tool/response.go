package tool

import (
	"github.com/gin-gonic/gin"
	"time"
)

const (
	REQUEST_SUCCESS = 200 //请求成功
	REQUEST_ERROR   = 400 //请求失败
)

//相应成功结构体
func ResponseSuccess(c *gin.Context, code int, msg string, result interface{}) {
	t := time.Now().Format("2006-1-2 15:04:05")
	c.JSON(200, gin.H{
		"code":   code,
		"msg":    msg,
		"result": result,
		"time":   t,
	})
}

//相应失败结构体
func ResponseFatal(c *gin.Context, code int, msg string, result interface{}) {
	t := time.Now().Format("2006-1-2 15:04:05")
	c.JSON(200, gin.H{
		"code":   code,
		"msg":    msg,
		"result": result,
		"time":   t,
	})
}
