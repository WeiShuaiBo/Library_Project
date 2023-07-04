package logic

import "github.com/gin-gonic/gin"

// Borrow godoc
//
// @Summary		借阅图书
// @Description 借阅指定书名的图书
// @Tags user
// @Router 	/user/book/{book_name} [POST]
func Borrow(c *gin.Context) {

}

// GiveBack godoc
//
// @Summary		归还图书
// @Description 归还指定书名的图书
// @Tags user
// @Router 		/user/book/{book_name} [GET]
func GiveBack(c *gin.Context) {

}

// GetUser godoc
//
// @Summary 查看用户信息
// @Description 查看用户信息
// @Tags user
// @Router /user [GET]
func GetUser(c *gin.Context) {

}

// PutUser godoc
//
// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags user
// @Router /user [PUT]
func PutUser(c *gin.Context) {

}

// GetMyBook godoc
//
// @Summary 查看自己的借阅记录
// @Description 查看自己的借阅记录
// @Tags user
// @Router /user/book [GET]
func GetMyBook(c *gin.Context) {

}
