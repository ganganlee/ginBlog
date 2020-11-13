package middle

import (
	. "blogV2.zozoo.net/tool"
	"github.com/gin-gonic/gin"
)

//验证用户登陆中间件
func LoginMiddleWare(context *gin.Context){
	authorization := context.GetHeader("Authorization")

	accessToken := AccessToken{
		Token: authorization,
	}

	secret, err := accessToken.ValidateToken()
	if err != nil {
		ResponseFatal(context,REQUEST_ERROR,"token验证失败",err.Error())
		context.Abort()
	}else if secret == "" {
		ResponseFatal(context,REQUEST_ERROR,"secret is empty","")
		context.Abort()
	}

	//设置用户secret
	context.Set("userSecret",secret)

	context.Next()
	return
}

