package controller

import (
	"blogV2.zozoo.net/dao"
	"blogV2.zozoo.net/service"
	"github.com/gin-gonic/gin"
	. "blogV2.zozoo.net/tool"
	"github.com/go-playground/validator/v10"
)

type ArticleController struct {
}

//发布文章
func (a *ArticleController) CreateArticle(c *gin.Context) {

	//获取用户信息
	userController := new(UserController)
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, "用户登陆失败", err.Error())
		return
	}

	//获取文章信息
	articleRequest := new(service.ArticleRequest)
	if err := c.ShouldBindJSON(articleRequest); err != nil {
		ResponseFatal(c, REQUEST_ERROR, "参数错误", GetValidateErr(err.(validator.ValidationErrors)))
		return
	}

	articleService := new(service.ArticleService)
	insertId, err := articleService.CreateArticle(articleRequest, user.Id)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, "文章添加失败", err.Error())
		return
	}

	//文章添加成功，启动携程循环添加到文章标签表中
	tagArticleDao := new(dao.TagArticleDao)
	for _, val := range articleRequest.Tag {
		Wg.Add(1)
		go tagArticleDao.Insert(val, insertId)
	}

	//等待携程同步
	Wg.Wait()

	ResponseSuccess(c, REQUEST_SUCCESS, "ok", insertId)
}

//获取文章详情
func (a *ArticleController) Info(c *gin.Context) {
	id := c.Param("id")
	articleService := new(service.ArticleService)
	info, err := articleService.Info(id)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, "获取文章失败", err)
		return
	}
	ResponseSuccess(c, REQUEST_SUCCESS, "ok", info)
}

//删除文章
func (a *ArticleController) DelArticle(c *gin.Context) {
	id := c.Param("id")
	articleService := new(service.ArticleService)
	err := articleService.DelArticle(id)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, "删除文章", err)
		return
	}

	ResponseSuccess(c, REQUEST_SUCCESS, "ok", nil)
}

//获取文章列表
func (a *ArticleController) List(c *gin.Context) {
	requestList := new(service.ArticleList)

	//获取参数
	if err := c.ShouldBindJSON(requestList); err != nil {
		ResponseFatal(c, REQUEST_ERROR, "获取参数失败", err)
		return
	}

	//查询数据
	articleService := new(service.ArticleService)
	list, total, err := articleService.List(requestList)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, "error", err)
		return
	}

	//返回结果
	ResponseSuccess(c, REQUEST_SUCCESS, "ok", gin.H{
		"total": total,
		"list":  list,
	})
}

//根据创建月份获取数据
func (a *ArticleController) ListByMonth(c *gin.Context) {
	articleService := new(service.ArticleService)
	list, groups, err := articleService.ListByMonth()
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, "获取数据失败", err)
		return
	}

	ResponseSuccess(c, REQUEST_SUCCESS, "ok", gin.H{
		"list":   list,
		"groups": groups,
	})
}
