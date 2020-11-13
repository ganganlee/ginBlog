package dao

import (
	"blogV2.zozoo.net/model"
	"blogV2.zozoo.net/tool"
)

type (
	UserDao struct {

	}
)

//插入数据
func (u *UserDao)Insert(user *model.User) error {
	orm := tool.Orm
	_, err := orm.InsertOne(user)
	if err != nil {
		return err
	}

	return nil
}

//用户登录
func (u *UserDao)Login(user *model.User) error  {
	orm := tool.Orm
	_, err := orm.Where("username = ? and password = ?", user.Username,user.Password).Get(user)
	return err
}

//根据secret查找用户
func (u *UserDao)QueryBySecret(user *model.User) error  {
	orm := tool.Orm
	_, err := orm.Where("secret = ?", user.Secret).Get(user)
	return err
}