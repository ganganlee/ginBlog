package service

import (
	"blogV2.zozoo.net/dao"
	"strconv"
)

type (
	TagArticleService struct {

	}
)

func (t *TagArticleService)List(id string,page string) (*[]dao.TagArticleDao,int64,error) {

	//tagId转int64
	atoi, err := strconv.Atoi(id)
	if err != nil {
		return nil,0,err
	}
	tagId := int64(atoi)

	//页数转int
	p,err := strconv.Atoi(page)

	//获取数据
	tagArticleDao := new(dao.TagArticleDao)
	list,count, err := tagArticleDao.List(tagId, p)
	if err != nil {
		return nil,0,err
	}

	return list,count,nil
}