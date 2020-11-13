package dao

import (
	"blogV2.zozoo.net/model"
	"blogV2.zozoo.net/tool"
	"errors"
)

type TagDao struct {
}

//根据名字查找tag
func (t *TagDao) RequestByName(name string, userId int64) ([]model.Tag, error) {
	orm := tool.Orm

	res := make([]model.Tag,0)
	err := orm.Where(" user_id = ? and name like ?", userId,"%"+name+"%").Find(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//根据名字添加tag
func (t *TagDao) Create(tag *model.Tag) error {
	orm := tool.Orm

	n, err := orm.Insert(tag)
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("添加失败，数据库繁忙")
	}

	return nil
}

//根据用户名查询用户tag列表
func (t *TagDao) List(userId int64)(*[]model.Tag,error)  {
	orm := tool.Orm

	tags := make([]model.Tag,0)
	if err := orm.Where("user_id = ?", userId).Find(&tags);err != nil {
		return nil, err
	}

	return &tags,nil
}
