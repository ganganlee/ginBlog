package service

import (
	"blogV2.zozoo.net/dao"
	"blogV2.zozoo.net/model"
	"blogV2.zozoo.net/tool"
	"crypto/md5"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

type (
	UserService struct {
	}

	//用户登陆结构体
	LoginUser struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	//登陆成功响应结构体
	LoginResponse struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}

	//用户注册请求结构体
	RegisterUser struct {
		LoginUser
		Email  string `json:"email" binding:"required"`
		Avatar string `json:"avatar" binding:"required"`
	}
)

//用户注册
func (u *UserService) Register(register *RegisterUser) error {

	//组织数据
	now := time.Now()
	user := &model.User{
		Username:   register.Username,
		Password:   fmt.Sprintf("%x", md5.Sum([]byte(register.Password))),
		Email:      register.Email,
		Avatar:     register.Avatar,
		Secret:     uuid.NewV4().String(),
		CreateTime: time.Unix(now.Unix(), 0),
		UpdateTime: time.Unix(now.Unix(), 0),
	}

	//调用dao曾写入数据库
	userDao := new(dao.UserDao)
	return userDao.Insert(user)
}

//用户登录
func (u *UserService) Login(login *LoginUser) (*LoginResponse, error) {
	//获取数据库模型
	dao := new(dao.UserDao)

	user := new(model.User)
	user.Username = login.Username
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(login.Password)))

	//查找用户书否存在
	if err := dao.Login(user); err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, errors.New("用户名或密码错误")
	}

	//生成jwtToken
	jwtToken := new(tool.AccessToken)
	jwtToken.Secret = user.Secret
	jwtToken.Expire = time.Now().Add(24 * time.Hour).Unix()
	if err := jwtToken.GenerateToken(); err != nil {
		return nil, err
	}

	loginResponse := new(LoginResponse)
	loginResponse.Username = user.Username
	loginResponse.Token = jwtToken.Token
	loginResponse.Avatar = user.Avatar
	loginResponse.Email = user.Email
	return loginResponse, nil
}

//查找用户
func (u *UserService) FindBySecret(secret string) (*model.User, error) {

	//获取用户信息
	userDao := new(dao.UserDao)

	user := new(model.User)
	user.Secret = secret
	if err := userDao.QueryBySecret(user); err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, errors.New("用户不存在")
	}

	return user, nil

	return nil, nil
}
