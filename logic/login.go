package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/model"
	"library/tools"
	"time"

	"net/http"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	Name string `form:"name" binding:"required"`
	Pwd  string `form:"pwd" binding:"required"`
}

type RUser struct {
	Name string `form:"name" binding:"required"`
	Pwd  string `form:"pwd" binding:"required"`
	Rpwd string `form:"rpwd" binding:"required"`
	Tel  string `form:"tel" binding:"required"`
}

// Login godoc
//
//	@Summary		用户登录
//	@Description	会执行用户登录操作
//	@Tags			Comm
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	true	"用户名"	Example(zhanggsan)
//	@Param			pwd		query		string	true	"密码"	Example(123)
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/login [POST]
func Login(c *gin.Context) {
	var user User
	n := c.Query("name")
	fmt.Printf("name:%s\n", n)
	if err := c.ShouldBind(&user); err != nil {
		fmt.Printf("绑定错误", err.Error())
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息校验错误",
			Data:    struct{}{},
		})
		return
	}

	dbUser := model.GetUser(user.Name, user.Pwd)
	if dbUser.Id <= 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息登录失败",
			Data:    struct{}{},
		})
		return
	}
	id := string(dbUser.Id)
	c.SetCookie("id", id, 0, "/", "localhost", false, true)
	//下发token
	a, r, err := tools.Token.GetToken(dbUser.Id, dbUser.Name)
	fmt.Printf("atoken:%s\n", a)
	fmt.Printf("rtoken:%s\n", r)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "获取token失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "登录成功",
		Data: Token{
			AccessToken:  a,
			RefreshToken: r,
		},
	})
	return
}

// Register godoc
//
//	@Summary		用户注册
//	@Description	会执行用户注册操作
//	@Tags			Comm
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	true	"用户名"
//	@Param			pwd		query		string	true	"密码"
//	@Param			tel		query		string	true	"手机号"
//	@Response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/register [POST]
func Register(c *gin.Context) {
	ruser := &model.User{}
	if err := c.ShouldBind(&ruser); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "注册用户信息格式校验错误",
			Data:    struct{}{},
		})
		return
	}
	ruser.CreatedTime = time.Now()
	if model.CheckGetUserExist(ruser.Name) == true {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户名重复创建失败",
			Data: struct {
			}{},
		})
		return
	}
	if err := model.RegisterUser(ruser); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户创建失败",
			Data: struct {
			}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "注册成功",
		Data: struct {
		}{},
	})
	return
}

// AdminLogin godoc
//
//	@Summary		管理员登录
//	@Description	会执行管理员登录操作
//	@Tags			Comm
//	@Accept			json
//	@Produce		json
//	@Param			name	query		string	true	"用户名"
//	@Param			pwd		query		string	true	"密码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/adminLogin [POST]
func AdminLogin(c *gin.Context) {
	var user User
	n := c.Query("name")
	fmt.Printf("name:%s\n", n)
	if err := c.ShouldBind(&user); err != nil {
		fmt.Printf("绑定错误", err.Error())
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "管理员信息校验错误",
			Data:    struct{}{},
		})
		return
	}

	dbUser := model.GetAdmin(user.Name, user.Pwd)
	if dbUser.Id <= 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "管理员信息登录失败",
			Data:    struct{}{},
		})
		return
	}
	id := string(dbUser.Id)
	c.SetCookie("id", id, 0, "/", "localhost", false, true)
	//下发token
	a, r, err := tools.Token.AdminGetToken(dbUser.Id, dbUser.Name)
	fmt.Printf("atoken:%s\n", a)
	fmt.Printf("rtoken:%s\n", r)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "管理获取token失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "管理员登录成功",
		Data: Token{
			AccessToken:  a,
			RefreshToken: r,
		},
	})
	return
}

//func Logout(c *gin.Context) {
//	_ = model.FlushSession(c)
//	c.JSON(http.StatusAccepted, tools.HttpCode{
//		Code: tools.OK,
//		Data: struct{}{},
//	})
//}
