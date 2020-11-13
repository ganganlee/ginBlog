package decorator

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//函数装饰器
func CacheDecorator(h gin.HandlerFunc,query string,format string) gin.HandlerFunc  {
	return func(c *gin.Context) {
		param := c.Param(query)
		key := fmt.Sprintf(format,param)
		fmt.Println(key)
		h(c)
	}
}