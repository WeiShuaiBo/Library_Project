package logic

import (
	"LibraryTest/dao"
	"LibraryTest/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User struct {
	Name string `json:"name" form:"name"`
	Pwd  string `json:"pwd" form:"pwd"`
}
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login godoc
//
// @Summary 用户登录
// @description 用户登录
// @Tags login
// @Router /log/login [POST]
func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
			Data:    struct{}{},
		})
		return
	}
	dbUser := dao.CheckUser(user.Name, user.Pwd)
	if dbUser.Id <= 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
			Data:    struct{}{},
		})
		return
	}
	a, r, err := tools.Token.GetToken(dbUser.Id, dbUser.Name)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "Token令牌生成失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "登录成功，正在跳转~",
		Data: Token{
			AccessToken:  a,
			RefreshToken: r,
		},
	})
	return
}

// Logout godoc
//
// @Summary 用户退出登录
// @description 用户退出登录
// @Tags login
// @Router /log/logout [GET]
func Logout(c *gin.Context) {
	_ = dao.FlushSession
	c.JSON(http.StatusUnauthorized, tools.HttpCode{
		Code: tools.OK,
		Data: struct{}{},
	})
	return
}

// Register godoc
//
// @Summary 用户注册
// @Description 用户注册
// @Tags login
// @Accept json
// @Produce json
// @Param name formData string false "用户名"
// @Param pwd formData string false "密码"
// @Param type formData string false "0:普通用户，1:超级用户"
// @response		200,400,500	{object}	tools.HttpCode{data=dao.User}
// @Router /log/register [POST]
func Register(c *gin.Context) {
	user := dao.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DOErr,
			Message: "用户注册信息绑定失败",
			Data:    struct{}{},
		})
		return
	}
	user.CreateTime = time.Now()
	fmt.Println(user)
	err := dao.Register(user)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DOErr,
			Message: "注册用户失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "用户注册成功，请登录~",
		Data:    struct{}{},
	})
	return
}
