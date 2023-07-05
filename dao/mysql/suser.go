package mysql

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create
// @Summary 测试创建用户
// @Description  用于生成用户
// @Tags 管理员
// Accept json
// Produce json
// @Router /api/v1/users [post]
func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "用户创建模块",
	})
}

// Delete
// @Summary 测试用户删除
// @Description  用于删除用户信息
// @Tags 管理员
// Accept json
// Produce json
// @Router /api/v1/users/:id [delete]
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "用户删除模块",
	})
}

// GetUsers
// @Summary 测试所以用户获取
// @Description  用于获取所有用户信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Router /api/v1/users [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "查询所有用户模块",
	})
}

// BCreate
// @Summary 创建图书
// @Description 创建图书模块
// @Tags 管理员
// @Accept json
// @Produce json
// @Router /api/v1/book [post]
func BCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "创建图书",
	})
}

// BDeleteSome
// @Summary 删除图书
// @Description 删除图书模块
// @Tags 管理员
// @Accept json
// @Produce json
// @Router /api/v1/book/:id [delete]
func BDeleteSome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "删除图书",
	})
}

// BorrowHistoryAll
// @Summary 借阅图书历史
// @Description 借阅图书历史模块
// @Tags 管理员
// @Accept json
// @Produce json
// @Router /api/v1/borrowHistory [get]
func BorrowHistoryAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "所有借阅历史",
	})
}

// BorrowHistorySome
// @Summary 借阅图书部分信息
// @Description 借阅图书部分信息
// @Tags 管理员
// @Accept json
// @Produce json
// @Router /api/v1/borrowHistory/:id [get]
func BorrowHistorySome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "模糊查询借阅历史",
	})
}
