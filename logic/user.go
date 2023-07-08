package logic

import (
	"LibraryTest/dao"
	"LibraryTest/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Borrow godoc
//
// @Summary		借阅图书
// @Description 借阅指定书名的图书
// @Tags user
// @Param Debug header string false "Debug header" default(123)
// @Param book_id path string false "图书ID"
// @response 200,500 {object} tools.HttpCode{}
// @Router 	/user/book/{book_id} [POST]
func Borrow(c *gin.Context) {
	//idAny, _ := c.Get("userId")
	//id := idAny.(int64)
	var id int64
	id = 1
	bookIdStr := c.Param("book_id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	ok := dao.Borrow(id, bookId)
	if !ok {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DOErr,
			Message: "借阅图书失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "图书借阅成功",
		Data:    struct{}{},
	})
	return
}

// GiveBack godoc
//
// @Summary		归还图书
// @Description 归还指定书名的图书
// @Tags user
// @Param Debug header string false "Debug" default(123)
// @Param book_id path string false "图书ID"
// @response 200,500 {object} tools.HttpCode{}
// @Router 		/user/book/{book_id} [GET]
func GiveBack(c *gin.Context) {
	//idAny, _ := c.Get("userId")
	//id := idAny.(int64)
	var id int64
	id = 1
	bookIdStr := c.Param("book_id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	ok := dao.GiveBack(id, bookId)
	if !ok {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DOErr,
			Message: "借阅归还失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "图书归还成功",
		Data:    struct{}{},
	})
	return
}

// GetUser godoc
//
// @Summary 查看用户信息
// @Description 查看用户信息
// @Tags user
// @Param Debug header string false "Debug header" default(123)
// @response 200,500 {object} tools.HttpCode{}
// @Router /user [GET]
func GetUser(c *gin.Context) {
	//id,err := c.Get("userId")
	//if err == false {
	//	c.JSON(http.StatusOK, tools.HttpCode{
	//		Code:    tools.UnLoginErr,
	//		Message: "用户还未登录",
	//		Data:    struct{}{},
	//	})
	//	return
	//}
	id := 1
	user := dao.GetUser(id)
	if user.Id < 1 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DOErr,
			Message: "用户信息查询失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "用户信息查询成功",
		Data:    user,
	})
	return
}

// PutUser godoc
//
// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags user
// @Param Debug header string false "Debug header" default(123)
// @Param name formData string false "用户名"
// @Param pwd formData string false "密码"
// @response 200,500 {object} tools.HttpCode
// @Router /user [PUT]
func PutUser(c *gin.Context) {
	//idAny, _ := c.Get("userId")
	//id := idAny.(int64)
	var id int64
	id = 1
	var user dao.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DOErr,
			Message: "用户信息绑定失败",
		})
		return
	}
	user.Id = id
	err := dao.PutUser(user)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DOErr,
			Message: "用户信息修改失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "用户信息修改成功",
		Data:    struct{}{},
	})
	return
}

// GetMyBook godoc
//
// @Summary 查看自己的借阅记录
// @Description 查看自己的借阅记录
// @Tags user
// @Param Debug header string false "Debug header" default(123)
// @Param page query string false "页数"
// @response 200,500 {object} tools.HttpCode{[]dao.book}
// @Router /user/book [GET]
func GetMyBook(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.ParseInt(pageStr, 10, 64)
	//idAny, _ := c.Get("userId")
	//id := idAny.(int64)
	var id int64
	id = 1
	recode := dao.GetMyBook(id, page)
	if len(recode) == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UnFoundErr,
			Message: "用户个人借阅记录查询失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "用户借阅信息查询成功",
		Data:    recode,
	})
	return

}
