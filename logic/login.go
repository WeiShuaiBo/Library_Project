package logic

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" form:"name"`
	Pwd  string `json:"pwd" form:"pwd"`
}

// Login godoc
//
// @Summary 用户登录
// @description 用户登录
// @Tags login
// @Router /login [POST]
func Login(c *gin.Context) {

}

// Logout godoc
//
// @Summary 用户退出登录
// @description 用户退出登录
// @Tags login
// @Router /logout [GET]
func Logout(c *gin.Context) {

}

// Register godoc
//
// @Summary 用户注册
// @description 用户注册
// @Tags login
// @Router /register [POST]
func Register(c *gin.Context) {

}
