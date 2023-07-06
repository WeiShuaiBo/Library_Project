package logic

import (
	"LibraryTest/dao"
	"LibraryTest/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GetAllUser godoc
//
// @Summary 查看所有用户信息
// @Description 查看所有用户信息
// @Tags manager
// @Router /manager/user [GET]
func GetAllUser(c *gin.Context) {
	user := dao.GetAllUser()
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
// @Router /manager/user/{id} [GET]
func GetUserBorrow(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	userBook := dao.GetUserBorrow(id)
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

// GetBookBorrow godoc
//
// @Summary 查看指定图书的借阅记录
// @Description 查看指定图书的借阅记录
// @Tags manager
// @Router /manager/{book_name} [GET]
func GetBookBorrow(c *gin.Context) {

}

// GetALlBookBorrow godoc
//
// @Summary 查看所有借阅记录
// @Description 查看所有借阅记录
// @Tags manager
// @Router /manager [GET]
func GetALlBookBorrow(c *gin.Context) {
	userBook := dao.GetALlBookBorrow()
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
	book.CreateTime = time.Now()
	fmt.Println(book)
	err := dao.AddBook(book)
	if err != nil {
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
		Data:    struct{}{},
	})
	return
}

// DeleteBook godoc
//
// @Summary		删除图书信息
// @Description 管理员删除指定书名的图书信息
// @Tags manager
// @Router 		/manager/book/{book_name} [DELETE]
func DeleteBook(c *gin.Context) {
	bookName := c.Param("book_name")
	fmt.Println(bookName)
	ok := dao.DeleteBook(bookName)
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
