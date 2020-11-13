package controller

import (
	"blogV2.zozoo.net/service"
	. "blogV2.zozoo.net/tool"
	"github.com/gin-gonic/gin"
)

type TagController struct {

}

//根据名字查找标签
func (t *TagController)FindByName(c *gin.Context)  {

	//获取用户信息
	userController := new(UserController)
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, "用户登陆失败", err.Error())
		return
	}

	//获取请求参数
	requestTag := new(service.RequestTag)

	name, ok := c.GetQuery("name")
	if !ok || name == "" {
		ResponseFatal(c,REQUEST_ERROR,"请求参数错误1",nil)
		return
	}
	requestTag.Name = name

	//前往数据库查找
	tagService := new(service.TagService)
	tags, err := tagService.FindByName(requestTag,user.Id)
	if err != nil {
		ResponseFatal(c,REQUEST_ERROR,"请求参数错误2",err)
		return
	}

	ResponseSuccess(c,REQUEST_SUCCESS,"ok",tags)
}

//添加tag
func (t *TagController)Create(c *gin.Context)  {

	//获取用户信息
	userController := new(UserController)
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, "用户登陆失败", err.Error())
		return
	}

	//获取请求从参数
	requestTag := new(service.RequestTag)
	if err := c.ShouldBindJSON(requestTag);err != nil {
		ResponseFatal(c,REQUEST_ERROR,"请求参数错误",err)
		return
	}

	//添加数据
	tagService := new(service.TagService)
	tag, err := tagService.Create(requestTag,user.Id)
	if err != nil {
		ResponseFatal(c,REQUEST_ERROR,"tag已存在",err)
		return
	}

	ResponseSuccess(c,REQUEST_SUCCESS,"ok",tag)
}

//获取用户tag列表
func (t *TagController)List(c *gin.Context)  {
	//获取用户信息
	userController := new(UserController)
	user, err := userController.GetUserInfo(c)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, "用户登陆失败", err.Error())
		return
	}

	//添加数据
	tagService := new(service.TagService)
	tags, err := tagService.List(user.Id)
	if err != nil {
		ResponseFatal(c,REQUEST_ERROR,"请求参数错误",err)
		return
	}

	ResponseSuccess(c,REQUEST_SUCCESS,"ok",tags)
}
