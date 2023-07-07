package system

import (
	"Library_Project/model/system/request"
	"Library_Project/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SysBookController struct {
}

// Books 书
// @Tags public
// @Summary 书
// @Accept json
// @Produce json
// @Param data body request.BookPage true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/public/books [get]
func (s *SysBookController) Books(c *gin.Context) {
	bookOff := &request.BookPage{}
	err := c.ShouldBindJSON(&bookOff)

	if err != nil {
		return
	}

	books := service.ServiceApp.SystemServiceGroup.Books(bookOff.Page, bookOff.Limit)
	c.JSON(200, gin.H{
		"code": 200,
		"date": books,
		"mgs":  "书籍",
	})
}

// Search 搜索书
// @Tags public
// @Summary 搜索书
// @Accept json
// @Produce json
// @Param BookName path string true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/public/search [get]
func (s *SysBookController) Search(c *gin.Context) {

	value := c.Query("BookName")
	fmt.Println(value)
	books := service.ServiceApp.SystemServiceGroup.SearchBooks(value)
	c.JSON(200, gin.H{
		"code": 200,
		"date": books,
		"mgs":  "书籍",
	})
}

// Records 借阅记录
// @Tags user
// @Summary 借阅记录
// @Accept json
// @Produce json
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/records [get]
func (s *SysBookController) Records(c *gin.Context) {
	userid := c.GetInt("UserId")
	re := service.ServiceApp.SystemServiceGroup.RecordPer(userid)
	c.JSON(200, gin.H{
		"code": 200,
		"date": re,
		"msg":  "借阅记录",
	})
}
