package logic

import (
	"AgroWeather/app/model"
	"AgroWeather/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"time"

	"net/http"
)

type User struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Pwd     string `json:"pwd" form:"pwd" binding:"required"`
	Tel     string `json:"tel" form:"tel" binding:"required"`
	Captcha string `json:"captcha" form:"captcha" binding:"required"`
}

type User1 struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
	Pwd1 string `json:"pwd1" form:"pwd1" binding:"required"`
	Pwd2 string `json:"pwd2" form:"pwd2" binding:"required"`
}

type User2 struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login godoc
//
//	@Summary		用户登录
//	@Description	会执行用户登录操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			pwd		formData	string	true	"密码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/login [POST]
func Login(c *gin.Context) {

	var user User2
	err := c.ShouldBind(&user)

	fmt.Println(user)

	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
			Data:    struct{}{},
		})
		return
	}

	//TODO: 入参校验 和 SQL注入问题
	dbUser := model.GetUser(user.Name, user.Pwd)

	if dbUser.Id < 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
			Data:    struct{}{},
		})
		return
	}

	a, r, err := tools.Token.GetToken(dbUser.Id, dbUser.Name)
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

// Register godoc
//
//	@Summary		用户注册
//	@Description	会执行用户注册操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			pwd		formData	string	true	"密码"
//	@Param			tel		formData	string	true	"联系方式"
//	@Param			captcha		formData	string	true	"验证码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/register [POST]
func Register(c *gin.Context) {

	var user User
	err := c.ShouldBind(&user)
	fmt.Println(user)
	if user.Tel == "" || user.Captcha == "" {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.CaptchaNil,
			Message: "验证码不能为空",
			Data:    struct{}{},
		})
		return
	}

	pool := model.NewPool("192.168.89.128:6379")
	conn := pool.Get()
	defer conn.Close()

	// 从 Redis 中获取存储的验证码
	storedCode, _ := redis.String(conn.Do("GET", user.Tel))
	fmt.Println(storedCode)

	// 比较用户输入的验证码和存储的验证码
	if user.Captcha != storedCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "验证码错误",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
			Data:    struct{}{},
		})
		return
	}

	if model.FindUserName(user.Name) {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserRepeatErr,
			Message: "用户名已经存在",
		})
		return
	}

	if model.Register(user.Name, user.Pwd, user.Tel) {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "注册成功",
		})
	}

	// 验证通过，删除存储的验证码
	if _, err := conn.Do("DEL", user.Tel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "验证码错误",
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.UserInfoErr,
		Message: "注册失败",
	})
}

// Logout godoc
//
//	@Summary		用户退出
//	@Description	会执行用户退出操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@response		500,401	{object}	tools.HttpCode
//	@Router			/logout [get]
func Logout(c *gin.Context) {
	_ = model.FlushSession(c)
	//这里暂时先不改为 401，有些接口确实不需要登录态
	c.JSON(http.StatusUnauthorized, tools.HttpCode{
		Code: tools.OK,
		Data: struct{}{},
	})
	return
}

// GetValidateCode godoc
//
//	@Summary		获取手机验证码
//	@Description	用于注册
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			tel		query	string	true	"联系方式"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/getValidateCode [POST]
func GetValidateCode(c *gin.Context) {
	tel := c.Query("tel")
	fmt.Println(tel)

	if tel == "" {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.CaptchaNil,
			Message: "手机号不能为空",
			Data:    struct{}{},
		})
		return
	}

	// 生成 6 位数验证码
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%06v", rand.Intn(1000000))
	fmt.Println("code:", code)
	pool := model.NewPool("192.168.89.128:6379")
	conn := pool.Get()
	defer conn.Close()

	// 存储验证码到 Redis，过期时间为 5 分钟
	if _, err := conn.Do("SETEX", tel, 5*60, code); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.CaptchaNil,
			Message: "验证码发送失败",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"code":   code,
	})
}

// PutUserPwd godoc
//
//	@Summary		用户修改密码
//	@Description	会执行修改密码操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			pwd		formData	string	true	"密码"
//	@Param			pwd1		formData	string	true	"新密码"
//	@Param			pwd2		formData	string	true	"确认密码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/putUserPwd [put]
func PutUserPwd(c *gin.Context) {

	var user User1
	err := c.ShouldBind(&user)

	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "用户信息错误",
			Data:    struct{}{},
		})
		return
	}

	if user.Pwd1 != user.Pwd2 || user.Pwd1 == "" {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "两次密码不同",
			Data:    struct{}{},
		})
		return
	}

	user1 := model.FindUserPwdByName(user.Name)

	if user.Pwd != user1.Pwd {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "密码错误",
			Data:    struct{}{},
		})
		return
	}

	ret := &model.User{
		Name: user.Name,
		Pwd:  user.Pwd1,
	}

	err1 := model.UpdateUserPwd(ret)
	if err1 != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "修改失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "修改成功，整在跳转~",
	})
	return
}
