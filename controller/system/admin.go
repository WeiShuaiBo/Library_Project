package system

import (
	"Library_Project/global"
	"Library_Project/model/system"
	"Library_Project/model/system/request"
	"Library_Project/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

// AddBook 添加图书
// @Tags admin
// @Summary 添加图书
// @Accept json
// @Produce json
// @Param data body system.Book true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/addbook [post]
func (s *AdminController) AddBook(c *gin.Context) {
	books := []system.Book{}
	err := c.ShouldBindJSON(&books)
	if err != nil {
		global.FAST_LOG.Error("添加书籍出错：" + err.Error())
		return
	}
	global.FAST_DB.Create(&books)
}

// DelBook 删除图书
// @Tags admin
// @Summary 删除图书
// @Accept json
// @Produce json
// @Param data body system.Book true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/delbook [post]
func (s *AdminController) DelBook(c *gin.Context) {
	books := []system.Book{}
	err := c.ShouldBindJSON(&books)
	if err != nil {
		global.FAST_LOG.Error("添加书籍出错：" + err.Error())
		return
	}
	global.FAST_DB.Delete(&books)
}

// SearchPerBook 搜索某人的借阅记录
// @Tags admin
// @Summary 搜索某人的借阅记录
// @Accept json
// @Produce json
// @Param data body request.Login true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/searchperbook [post]
func (s *AdminController) SearchPerBook(c *gin.Context) {
	u := &request.Login{}
	err := c.ShouldBindJSON(&u)
	fmt.Println(u)
	if err != nil {
		global.FAST_LOG.Error("搜索某人借阅记录：" + err.Error())
		return
	}
	re := service.ServiceApp.SystemServiceGroup.RecordBook(u.UserId)
	c.JSON(200, gin.H{
		"code": 200,
		"date": re,
		"msg":  "借阅记录",
	})
}

// SearchBook 搜索某书的借阅记录
// @Tags admin
// @Summary 搜索某书的借阅记录
// @Accept json
// @Produce json
// @Param data body request.ReqBook true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/searchbook [post]
func (s *AdminController) SearchBook(c *gin.Context) {
	reqBook := &request.ReqBook{}
	err := c.ShouldBindJSON(reqBook)
	fmt.Println(reqBook)
	if err != nil {
		global.FAST_LOG.Error("搜索书籍出错：" + err.Error())
		return
	}

	re := service.ServiceApp.SystemServiceGroup.RecordBook(reqBook.Id)
	c.JSON(200, gin.H{
		"code": 200,
		"date": re,
		"msg":  "被借记录",
	})
}
