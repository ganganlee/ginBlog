package dao

import (
	"blogV2.zozoo.net/model"
	"blogV2.zozoo.net/tool"
	"errors"
)

type TagArticleDao struct {
	model.Article    `xorm:"extends"`
}

//添加文章标签
func (t *TagArticleDao) Insert(tagId, articleId int64) {
	orm := tool.Orm

	tagArticle := &model.TagArticle{
		TagId:     tagId,
		ArticleId: articleId,
	}
	orm.InsertOne(tagArticle)
	tool.Wg.Done()
}

//根据tagId获取文章列表
func (t *TagArticleDao) List(tagId int64, page int) (*[]TagArticleDao,int64,error) {

	orm := tool.Orm

	//获取总条数
	tagArticle := new(model.TagArticle)
	count, err := orm.Where("tag_id = ?",tagId).Count(tagArticle)
	if err != nil{
		return nil, 0, err
	}

	if count == 0 {
		return nil, 0, errors.New("数据为空")
	}

	articles := make([]TagArticleDao, 0)
	err = orm.SQL("select article.* from tag_article left join article on tag_article.article_id = article.id where tag_article.tag_id = ?",tagId).Find(&articles)

	if err != nil {
		return nil,0,err
	}

	return &articles,count,nil
}
