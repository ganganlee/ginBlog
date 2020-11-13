package dao

import (
	"blogV2.zozoo.net/model"
	"blogV2.zozoo.net/tool"
	"errors"
)

type (
	ArticleDao struct {
	}

	ArticleList struct {
		model.Article `xorm:"extends"`
		Month         string `json:"month"`
	}
)

//添加文章
func (a *ArticleDao) Insert(article *model.Article) (int64, error) {
	orm := tool.Orm
	n, err := orm.InsertOne(article)
	if err != nil {
		return 0, err
	}

	if n == 0 {
		return 0, errors.New("数据库繁忙")
	}

	return article.Id, nil
}

//获取文章详情
func (a *ArticleDao) Get(id int64) (*model.Article, error) {
	article := new(model.Article)
	article.Id = id

	orm := tool.Orm
	ok, err := orm.Where("id = ?", article.Id).Get(article)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("文章不存在")
	}

	return article, nil
}

//获取文章列表
func (a *ArticleDao) List(page, pageSize int) (*[]model.Article, int64, error) {
	orm := tool.Orm

	//获取总条数
	article := new(model.Article)
	total, err := orm.Count(article)

	if err != nil {
		return nil, 0, err
	}

	//获取数据
	list := make([]model.Article, 0)
	if err := orm.Select("id,user_id,title,synopsis,tag_str,sort,view,comment,create_time").Limit(pageSize, (page-1)*page).Find(&list); err != nil {
		return nil, total, err
	}

	return &list, total, nil
}

//根据时间分组获取数据
func (a *ArticleDao) ListByMonth() (*[]ArticleList, error) {

	orm := tool.Orm

	articles := make([]ArticleList, 0)
	err := orm.SQL("SELECT id,user_id,title,synopsis,tag_str,sort,view,comment,create_time, DATE_FORMAT(create_time,'%m月，%Y') as month FROM  article ORDER BY id DESC").Find(&articles)
	if err != nil {
		return nil, err
	}
	return &articles, nil
}

//删除文章，根据id
func (a *ArticleDao) DelById(id int64) error {
	articleModel := new(model.Article)
	orm := tool.Orm
	_, err := orm.Where("id = ?", id).Delete(articleModel)
	if err != nil {
		return err
	}

	//删除标签文章表
	articleTag := new(model.TagArticle)
	_, err = orm.Where("article_id = ?", id).Delete(articleTag)
	if err != nil {
		return err
	}

	return nil
}
