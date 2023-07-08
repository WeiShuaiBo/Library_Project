package logic

import (
	"LibraryTest/dao"
	"LibraryTest/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	time2 "time"
)

// GetAllUser godoc
//
// @Summary 查看所有用户信息
// @Description 查看所有用户信息
// @Tags manager
// @Param Debug header string false "Debug header" default(123)
// @Param page query string false "页数"
// @response 200,500 {object} tools.HttpCode{[]dao.User}
// @Router /manager/user [GET]
func GetAllUser(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	user := dao.GetAllUser(page)
	if len(user) == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UnFoundErr,
			Message: "未找到全部用户信息",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "成功找到所有用户信息",
		Data:    user,
	})
}

// GetUserBorrow godoc
//
// @Summary 查看指定用户的借阅信息
// @Description 查看指定用户的借阅信息
// @Tags manager
// @Param Debug header string false "Debug header" default(123)
// @Param page query string false "页数"
// @Param id path string false "用户Id"
// @response 200,500 {object} tools.HttpCode{[]dao.UserBook}
// @Router /manager/user/{id} [GET]
func GetUserBorrow(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	userBook := dao.GetUserBorrow(id, page)
	if len(userBook) == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UnFoundErr,
			Message: "未找到指定用户的借阅记录",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "成功找到指定用户的借阅记录",
		Data:    userBook,
	})
	return
}

// GetALlBookBorrow godoc
//
// @Summary 查看所有借阅记录
// @Description 查看所有借阅记录
// @Tags manager
// @Param Debug header string false "Debug header" default(123)
// @Param page query string false "页数"
// @response 200,500 {object} tools.HttpCode{[]dao.UserBook}
// @Router /manager [GET]
func GetALlBookBorrow(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	userBook := dao.GetALlBookBorrow(page)
	if len(userBook) == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UnFoundErr,
			Message: "未找到全部图书的借阅信息",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "成功找到全部图书的借阅信息",
		Data:    userBook,
	})
	return
}

// AddBook godoc
//
// @Summary		添加图书
// @Description 管理员添加图书信息
// @Tags manager
// @Param id formData int false "图书ID"
// @Param book_name formData string false "书名"
// @Param author formData string false "作者"
// @Param publishing_house formData string false "出版社"
// @Param translator formData string false "译者"
// @Param publish_data formData string true "日期：2002-03-02"
// @Param pages formData int false "页数"
// @Param isbn formData string false "isbn编号"
// @Param price formData number false "价格"
// @Param brief_introduction formData string false "图书简介"
// @Param author_introduction formData string false "作者简介"
// @Param img_url formData string false "图片链接"
// @Param count formData int false "图书数量"
// @response 200,500 {object} tools.HttpCode{dao.Book}
// @Router 		/manager/book [POST]
func AddBook(c *gin.Context) {
	book := dao.Book{}
	if err := c.ShouldBind(&book); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DOErr,
			Message: "图书信息绑定失败",
			Data:    struct{}{},
		})
		return
	}
	timeStr, _ := c.GetPostForm("publish_data")
	fmt.Println(timeStr)
	layout := "2006-01-02"
	time, _ := time2.Parse(layout, timeStr)
	fmt.Println(time)
	book.PublishDate = time
	ok := dao.AddBook(book)
	if !ok {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DOErr,
			Message: "图书添加失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "图书添加成功",
		Data:    book,
	})
	return
}

// DeleteBook godoc
//
// @Summary		删除图书信息
// @Description 管理员删除指定书名的图书信息
// @Tags manager
// @Param Debug header string false "Debug header" default(123)
// @Param book_id path string false "书名"
// @response 200,500 {object} tools.HttpCode{}
// @Router 		/manager/book/{book_id} [DELETE]
func DeleteBook(c *gin.Context) {
	bookIdStr := c.Param("book_id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	ok := dao.DeleteBook(bookId)
	if !ok {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UnFoundErr,
			Message: "未能删除指定图书",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "成功删除指定图书",
		Data:    struct{}{},
	})
	return
}
