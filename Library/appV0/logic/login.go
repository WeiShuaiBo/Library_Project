package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"library/appV0/model"
	"library/appV0/tools"
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
//	@Summary		登录
//	@Description	会执行用户登录操作
//	@Tags			guest
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			pwd		formData	string	true	"密码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/guest/login [POST]
func Login(c *gin.Context) {

	user := User2{}
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

	fmt.Println(dbUser.Id)

	if dbUser.Id < 1 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "账号或者密码不正确",
			Data:    struct{}{},
		})
		return

	}

	a, r, err1 := tools.Token.GetToken(dbUser.Id, dbUser.Name)
	fmt.Printf("atoken:%s\n", a)
	fmt.Printf("rtoken:%s\n", r)

	if err1 != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "Token生成失败！错误信息：" + err1.Error(),
			Data:    struct{}{},
		})
		return
	}
	// 记录日志并使用zap.Xxx(key, val)记录相关字段
	zap.L().Debug("this is 登录 func", zap.String("user", user.Name), zap.String("pwd", user.Pwd))

	c.String(http.StatusOK, "hello 王志乾.com!")

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
//	@Summary		注册
//	@Description	会执行用户注册操作
//	@Tags			guest
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			pwd		formData	string	true	"密码"
//	@Param			tel		formData	string	true	"联系方式"
//	@Param			captcha		formData	string	true	"验证码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/guest/register [POST]
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

	// 从 Redis 中获取存储的验证码
	storedCode, _ := redis.String(model.RedisConn.Get().Do("GET", user.Tel))
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

	// 验证通过，删除存储的验证码
	if _, err := model.RedisConn.Get().Do("DEL", user.Tel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "验证码错误",
		})
		return
	}

	if model.Register(user.Name, user.Pwd, user.Tel) {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "注册成功",
		})
	}

}

// Logout godoc
//
//	@Summary		用户退出
//	@Description	会执行用户退出操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@response		500,401	{object}	tools.HttpCode
//	@Router			/user/logout [get]
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
//	@Tags			guest
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			tel		query	string	true	"联系方式"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/guest/getValidateCode [POST]
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

	// 存储验证码到 Redis，过期时间为 5 分钟
	if _, err := model.RedisConn.Get().Do("SETEX", tel, 5*60, code); err != nil {
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
//	@Summary		修改密码
//	@Description	会执行修改密码操作
//	@Tags			user/common
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			pwd		formData	string	true	"密码"
//	@Param			pwd1		formData	string	true	"新密码"
//	@Param			pwd2		formData	string	true	"确认密码"
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/user/common/putUserPwd [put]
func PutUserPwd(c *gin.Context) {

	var user User1

	err := c.ShouldBind(&user)

	name := ""

	n, exists := c.Get("name")

	if exists {
		name = n.(string)
	}

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

	if user.Pwd1 == user.Pwd {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "新密码和旧密码不能相同",
			Data:    struct{}{},
		})
		return
	}

	user1 := model.FindUserPwdByName(name)

	fmt.Println("user1:", user1)

	if user1.Name == "" {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "请先登录",
			Data:    struct{}{},
		})
		return
	}

	if user.Pwd != user1.Pwd {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "密码错误",
			Data:    struct{}{},
		})
		return
	}

	ret := &model.User{
		Name: name,
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
		Message: "修改成功",
	})
}
