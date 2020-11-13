package controller

import (
	"blogV2.zozoo.net/model"
	"blogV2.zozoo.net/service"
	. "blogV2.zozoo.net/tool"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	UserController struct {
	}
)

//用户注册
func (u *UserController) Register(c *gin.Context) {
	//获取请求参数
	registerUser := new(service.RegisterUser)
	err := c.BindJSON(registerUser)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, err.Error(), "")
		return
	}

	//调用service层处理注册逻辑
	userService := new(service.UserService)
	err = userService.Register(registerUser)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, err.Error(), "")
		return
	}

	ResponseSuccess(c, REQUEST_SUCCESS, "ok", nil)
	return
}

//用户登陆
func (u *UserController) Login(c *gin.Context) {
	login := new(service.LoginUser)

	//获取登陆参数
	if err := c.ShouldBindJSON(login); err != nil {
		ResponseFatal(c, REQUEST_ERROR, "登录失败", GetValidateErr(err.(validator.ValidationErrors)))
		return
	}

	//调用逻辑层处理业务逻辑
	userService := new(service.UserService)
	response, err := userService.Login(login)
	if err != nil {
		ResponseFatal(c, REQUEST_ERROR, err.Error(), "")
		return
	}

	ResponseSuccess(c, REQUEST_SUCCESS, "ok", response)
	return
}

//获取用户信息
func (u *UserController) GetUserInfo(c *gin.Context) (*model.User, error) {
	//获取用户信息
	val, _ := c.Get("userSecret")
	secret := val.(string)
	if secret == "" {
		return nil, errors.New("用户登陆失败")
	}

	userService := new(service.UserService)
	user, err := userService.FindBySecret(secret)
	if err != nil {
		return nil, err
	}

	return user, nil
}
