package logic

import (
	"Library-management/appv1/modle"
	"Library-management/appv1/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login godoc
//
//	@Summary		登录
//	@Description	会执行登录操作
//	@Tags			all
//	@Accept			mpfd
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			password		formData	string	true	"密码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/all/login [POST]
func Login(c *gin.Context) {
	var user User
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户输入信息错误",
			Data:    struct{}{},
		})
		return
	}

	dbUser := modle.GetUser(user.Name, modle.Md(user.Password))
	if dbUser.UserId <= 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
			Data:    struct{}{},
		})
		return
	}

	//下发Token
	a, r, err := tools.Token.GetToken(dbUser.UserId, dbUser.Name)
	fmt.Printf("atoken:%s\n", a)
	fmt.Printf("rtoken:%s\n", r)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "Token生成失败！错误信息：" + err.Error(),
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "登录成功，整在跳转~",
		Data: Token{
			AccessToken:  a,
			RefreshToken: r,
		},
	})
	return
}
