package route

import (
	"blogV2.zozoo.net/controller"
	"blogV2.zozoo.net/middle"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoute(gin *gin.Engine) {

	//支持跨域
	gin.Use(Cors())

	//实例化用户控制器
	userController := &controller.UserController{}

	//用户模块
	userGrout := gin.Group("/user")
	userGrout.POST("/register", userController.Register)
	userGrout.POST("/login", userController.Login)

	//文章模块
	article := new(controller.ArticleController)
	gin.GET("/article/:id", article.Info)                         //获取文章详情
	gin.POST("/article/:id", article.DelArticle)                         //删除文章
	gin.POST("/articles", article.List)                           //获取文章列表
	gin.POST("/articles/month", article.ListByMonth)              //根据月份获取文章详情
	articleGroup := gin.Group("/article", middle.LoginMiddleWare) //注册登陆中间件
	articleGroup.POST("/", article.CreateArticle)                 //创建文章

	//tag模块
	tag := new(controller.TagController)
	tagGroup := gin.Group("/tag", middle.LoginMiddleWare)
	tagGroup.POST("", tag.Create)    //添加tag
	tagGroup.GET("", tag.FindByName) //搜索用户tag
	tagGroup.GET("/list", tag.List)  //获取用户tag

	//标签文章模块
	tagArticle := new(controller.TagArticleController)
	gin.GET("/tagArticle/:id", tagArticle.List) //根据文章标签获取标签下的文章
}

//Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
