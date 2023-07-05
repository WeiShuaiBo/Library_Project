// @Author	zhangjiaozhu 2023/7/3 17:01:00
package api

import (
	"MyBook/common/Response"
	Email "MyBook/common/SendEmail"
	"MyBook/common/jwt"
	"MyBook/dao"
	"MyBook/models"
	redis "MyBook/models/DB/RedisDB"
	"MyBook/routers/midwares"
	"MyBook/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// UserRegister
// @Summary 用户注册
// @Tags  公共用户
// @Accept json
// @Produce json
// @Param data body models.UserRegister true "请求参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	// 获取注册请求参数，校验数据的有效性
	var userRegister models.UserRegister
	if err := c.ShouldBind(&userRegister); err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	if userRegister.UserName == "" || userRegister.Password == "" || userRegister.Code == "" || userRegister.Email == "" {
		Response.Error(c, "参数不能为空")
		return
	}
	// 校验验证码
	result, err := redis.RDB.Get(c, userRegister.Email).Result()
	if err != nil {
		Response.Error(c, "服务器繁忙")
		return
	}
	if result != userRegister.Code {
		Response.Error(c, "无效的验证码")
		return
	}
	// 注册用户
	if err := service.UserRegister(&userRegister); err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "注册成功", nil)
}

// UserLogin
// @Summary 用户登陆
// @Tags  公共用户
// @Accept json
// @Produce json
// @Param data body models.UserLogin true "请求参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var u models.UserLogin
	err := c.ShouldBindJSON(&u)
	if err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	if u.UserName == "" || u.Password == "" || u.Code == "" || u.Email == "" {
		Response.Error(c, "参数不能为空")
		return
	}
	// 校验验证码
	result, err := redis.RDB.Get(c, u.Email).Result()
	if err != nil {
		Response.Error(c, "服务器繁忙")
		return
	}
	if result != u.Code {
		Response.Error(c, "无效的验证码")
		return
	}
	// 用户登陆
	user, err := service.UserLogin(u)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	// 发送token
	aToken, rToken, _ := jwt.GenToken(user.UserId)
	Response.Success(c, "success", gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
		"userID":       user.UserId,
		"username":     user.UserName,
	})
}

// FindUserInfo
// @Summary 查看用户信息
// @Tags  用户
// @Accept json
// @Produce json
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/findUserInfo [get]
func FindUserInfo(c *gin.Context) {
	value, exists := c.Get(midwares.ContextUserIDKey)
	if !exists {
		Response.Error(c, "未登录")
		return
	}
	u := value.(uint64)
	user, err := service.FindUserById(u)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "success", user)
}

// UpdateUserInfo
// @Summary 更新用户信息
// @Tags  用户
// @Accept json
// @Produce json
// @Param data body models.UserUpdate true "请求参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/updateUserInfo [post]
func UpdateUserInfo(c *gin.Context) {

	type UserUpdate struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	var userUpdate UserUpdate
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	// 获取当前用户的id
	value, exists := c.Get(midwares.ContextUserIDKey)
	if !exists {
		Response.Error(c, "未登录")
		return
	}
	u := value.(uint64)
	user, err := dao.FindUserById(u)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	user.UserName = userUpdate.UserName
	user.Password = userUpdate.Password
	user.Email = userUpdate.Email
	err = service.UpdateUserInfo(user)
	if err != nil {
		Response.Error(c, err.Error())
	}
	Response.Success(c, "success", nil)
}

func RefreshTokenHandler(c *gin.Context) {
	rt := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		Response.Error(c, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		Response.Error(c, "Token格式不对")
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

func Logout(c *gin.Context) {
	Response.Success(c, "请自行清除token", nil)
}

// DeleteUser
// @Summary 删除用户
// @Tags  管理员
// @Accept json
// @Produce json
// @Param data body Req true "请求参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/findUserInfo [get]
func DeleteUser(c *gin.Context) {
	type Req struct {
		UserId string `json:"user_id"`
	}
	var req Req
	err := c.ShouldBindJSON(&req)
	if err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	err = service.DeleteUser(req.UserId)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "success", nil)
}

// GetAllUser
// @Summary 查看所有用户信息
// @Tags  管理员
// @Accept json
// @Produce json
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user/findUserInfo [get]
func GetAllUser(c *gin.Context) {
	user, err := service.GetAllUser()
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "success", user)
}

func SendEmail(c *gin.Context) {
	type Req struct {
		Mail string `json:"mail"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	checkCode := GetRand()
	// 将验证码存入redis,设置有效期
	redis.RDB.Set(c, req.Mail, checkCode, time.Second*300)
	// 将验证码发送到指定邮箱
	err := Email.SendEmail(req.Mail, checkCode)
	if err != nil {
		Response.Error(c, "验证码发送失败")
		return
	}
	Response.Success(c, "success", nil)
}

// 生成6位随机验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
