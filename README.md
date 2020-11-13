# ginBlog
>基于gin开发的博客API接口，使用jwt进行验证和登陆
## 使用方法
- git clone https://github.com/ganganlee/ginBlog.git
- go mod tidy
- 修改/config/config.json里面的配置文件
- 直接运行或者编译（编译参考：https://blog.csdn.net/weixin_44540711/article/details/109682597）
## 目录结构参考

```bash
/-config #配置文件
/-controller #控制器
/-dao #数据库操作层
/-decorator #装饰器层（还没实现）
/-middle #中间件
/-route #路由
/-service #逻辑层
/-tool #公共函数层
main.go 入口文件
```
## 路由地址参考

```bash
	//用户模块
	userGrout := gin.Group("/user")
	userGrout.POST("/register", userController.Register)
	userGrout.POST("/login", userController.Login)

	//文章模块
	article := new(controller.ArticleController)
	gin.GET("/article/:id", article.Info)                         //获取文章详情
	gin.POST("/article/:id", article.DelArticle)                  //删除文章
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
```
