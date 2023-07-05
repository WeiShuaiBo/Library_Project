package mysql

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login
// @Summary 登录接口
// @Description  登录
// @Tags 公开
// @Accept json
// @Produce json
// @Router /api/v1/public/login [post]
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "此处是用户登录模块",
	})
}

// Register
// @Summary 注册接口
// @Description  注册
// @Tags 公开
// @Accept json
// @Produce json
// @Router /api/v1/public/register [post]
func Register(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "此处是用户注册模块",
	})
}

// Logout
// @Summary 退出登录接口
// @Description 退出登录
// @Tags 公开
// @Accept json
// @Produce json
// @Router /api/v1/public/logout [post]
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "退出登录了",
	})
}

// BGetBooks
// @Summary 获取图书列表
// @Description 获取所有图书信息
// @Tags 公开
// @Accept json
// @Produce json
// @Router /api/v1/public/BGetBooks [post]
func BGetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "获取图书列表",
	})
}

// BGetBook
// @Summary 模糊查询图书
// @Description 模糊查询图书信息
// @Tags 公开
// @Accept json
// @Produce json
// @Router /api/v1/public/BGetBook [post]
func BGetBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "模糊查询图书信息",
	})
}
