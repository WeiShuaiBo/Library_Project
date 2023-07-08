package logic

import (
	"LibraryTest/dao"
	"LibraryTest/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetBook godoc
//
// @Summary		获取图书信息
// @Description 通过模糊查询查看图书信息学
// @Tags 		library
// @Accept json
// @Produce json
// @Param book_name path string true "书名"
// @Param page query string false "页数"
// @response 200,500 {object} tools.HttpCode{data=[]dao.Book}
// @Router 		/{book_name} [GET]
func GetBook(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	bookName := c.Param("book_name")
	book := dao.GetBook(bookName, page)
	if len(book) == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UnFoundErr,
			Message: "未找到指定图书",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "已找到指定图书",
		Data:    book,
	})
}

// GetAllBook godoc
//
// @Summary 	获取所有图书信息
// @Description 获取所有图书信息
// @Tags 		library
// @Param page query string false "需要查询的页数"
// @response 200,500 {object} tools.HttpCode{data=[]dao.Book}
// @Router 		/  [Get]
func GetAllBook(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	book := dao.GetAllBook(page)
	if len(book) == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UnFoundErr,
			Message: "全部图书信息查询失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "全部图书信息查询成功",
		Data:    book,
	})
}
