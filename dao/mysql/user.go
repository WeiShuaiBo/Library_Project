package mysql

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser
// @Summary 测试用户获取
// @Description  用于获取单个用户信息
// @Tags 普通用户
// Accept json
// Produce json
// @Router /api/v1/selfInformation [get]
func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "查询单个用户模块",
	})
}

// Update
// @Summary 测试用户更新
// @Description  用于更新用户信息
// @Tags 普通用户
// @Accept json
// @Produce json
// @Router /api/v1/updateInformation [put]
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "用户更新模块",
	})
}

// Borrow
// @Summary 借书模块
// @Description 用户借书
// @Tags 普通用户
// @Accept json
// @Produce json
// @Router /api/v1/borrow/:id [post]
func Borrow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "借书",
	})
}

// Return
// @Summary 还书模块
// @Description 用户还书
// @Tags 普通用户
// @Accept json
// @Produce json
// @Router /api/v1/return/:id [post]
func Return(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "还书",
	})
}

// GetSelfBorrowHistory
// @Summary 查看借阅历史
// @Description 查看个人借阅历史
// @Tags 普通用户
// @Accept json
// @Produce json
// @Router /api/v1/selfBorrowHistory [get]
func GetSelfBorrowHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "查阅历史信息",
	})
}
