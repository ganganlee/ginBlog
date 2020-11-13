package controller

import (
	"blogV2.zozoo.net/service"
	. "blogV2.zozoo.net/tool"
	"github.com/gin-gonic/gin"
)

type TagArticleController struct {
}

//根据标签id获取文章列表
func (t *TagArticleController) List(c *gin.Context) {
	//获取文章id
	id := c.Param("id")
	//获取page
	page := c.DefaultQuery("page","1")

	tagArticleService := new(service.TagArticleService)
	list,total, err := tagArticleService.List(id, page)
	if err != nil {
		ResponseFatal(c,REQUEST_ERROR,"error",err)
		return
	}
	ResponseSuccess(c, REQUEST_SUCCESS, "ok", gin.H{
		"list":list,
		"total":total,
	})
}
