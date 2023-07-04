package logic

import (
	"LibraryTest/dao"
	"LibraryTest/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBook godoc
//
// @Summary		获取图书信息
// @Description 通过模糊查询查看图书信息学
// @Tags 		library
// @Accept json
// @Produce json
// @Param book_name path string true "书名"
// @response 200,500 {object} tools.HttpCode{data=dao.Book}
// @Router 		/{book_name} [GET]
func GetBook(c *gin.Context) {
	bookName := c.Param("book_name")
	book := dao.GetBook(bookName)
	if book.Id <= 0 {
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
// @Router 		/  [Get]
func GetAllBook(c *gin.Context) {
	book := dao.GetAllBook()
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
