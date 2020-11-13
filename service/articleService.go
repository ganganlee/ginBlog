package service

import (
	"blogV2.zozoo.net/dao"
	"blogV2.zozoo.net/model"
	"errors"
	"strconv"
)

type (
	ArticleService struct {
	}

	//添加文章结构体
	ArticleRequest struct {
		Title    string   `json:"title" binding:"required"`    //标题
		Synopsis string   `json:"synopsis" binding:"required"` //简介
		Content  string   `json:"content" binding:"required"`  //内容
		Tag      []int64  `json:"tag" binding:"required"`      //tagId 切片
		TagStr   []string `json:"tag_str" binding:"required"`  //tagStr 切片
	}

	//获取文章列表
	ArticleList struct {
		Page     int `json:"page" binding:"required"`
		PageSize int `json:"page_size" binding:"required"`
	}
)

//添加文章
func (a *ArticleService) CreateArticle(art *ArticleRequest, userId int64) (int64, error) {
	article := &model.Article{
		UserId:   userId,
		Title:    art.Title,
		Synopsis: art.Synopsis,
		Content:  art.Content,
		Tag:      art.Tag,
		TagStr:   art.TagStr,
	}

	articleDao := new(dao.ArticleDao)
	insertId, err := articleDao.Insert(article)
	if err != nil {
		return 0, err
	}

	return insertId, nil
}

//获取文章详情
func (a *ArticleService) Info(id string) (*model.Article, error) {

	articleDao := new(dao.ArticleDao)

	//将字符串id转换为int64
	tmpId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("参数解析失败")
	}
	articleId := int64(tmpId)

	//获取文章详情
	article, err := articleDao.Get(articleId)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (c *ArticleService) DelArticle(id string) error {
	articleDao := new(dao.ArticleDao)

	//将字符串id转换为int64
	tmpId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("参数解析失败")
	}
	articleId := int64(tmpId)

	//删除文章
	if err := articleDao.DelById(articleId); err != nil {
		return err
	}

	return nil
}

//获取文章列表
func (a *ArticleService) List(res *ArticleList) (*[]model.Article, int64, error) {
	//条件判断
	if res.Page < 1 || res.PageSize < 1 {
		return nil, 0, errors.New("参数错误")
	}

	articleDao := new(dao.ArticleDao)
	list, total, err := articleDao.List(res.Page, res.PageSize)
	if err != nil {
		return nil, total, err
	}
	return list, total, nil
}

//根据时间分组获取书
func (a *ArticleService) ListByMonth() (*[]dao.ArticleList, *[]string, error) {
	articleDao := new(dao.ArticleDao)
	articles, err := articleDao.ListByMonth()
	if err != nil {
		return nil, nil, err
	}

	//获取所有文章分组
	groups := make([]string, 0)
	months := make(map[string]string, 0)
	for key, val := range *articles {
		months[val.Month] = val.Month
		(*articles)[key].Content = ""
	}
	for _, val := range months {
		groups = append(groups, val)
	}

	return articles, &groups, nil
}
