package service

import (
	"blogV2.zozoo.net/dao"
	"blogV2.zozoo.net/model"
)

type (
	TagService struct {

	}

	//根据名字查找tag
	RequestTag struct {
		Name string `json:"name" binding:"required"`
	}
)

//根据名字查找tag
func (t *TagService)FindByName(res *RequestTag,userId int64)([]model.Tag,error)  {
	tagDao := new(dao.TagDao)

	return tagDao.RequestByName(res.Name,userId)
}

//添加tag
func (t *TagService)Create(res *RequestTag,userId int64)(*model.Tag,error)  {
	tagDao := new(dao.TagDao)

	tag := new(model.Tag)
	tag.Name = res.Name
	tag.UserId = userId
	tag.Status = 1

	if err := tagDao.Create(tag);err != nil {
		return nil,err
	}

	return tag,nil
}

//获取用户tag列表
func (t *TagService)List(userId int64) (*map[int64]string,error) {
	tag := new(dao.TagDao)
	tags, err := tag.List(userId)
	if err != nil {
		return nil,err
	}

	list := make(map[int64]string,0)
	for _,val := range *tags {
		list[val.Id] = val.Name
	}

	return &list,nil
}