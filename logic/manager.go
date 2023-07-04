package logic

import "github.com/gin-gonic/gin"

// GetAllUser godoc
//
// @Summary 查看所有用户信息
// @Description 查看所有用户信息
// @Tags manager
// @Router /manager/user [GET]
func GetAllUser(c *gin.Context) {

}

// GetUserBorrow godoc
//
// @Summary 查看指定用户的借阅信息
// @Description 查看指定用户的借阅信息
// @Tags manager
// @Router /manager/user/{id} [GET]
func GetUserBorrow(c *gin.Context) {

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

}

// AddBook godoc
//
// @Summary		添加图书
// @Description 管理员添加图书信息
// @Tags manager
// @Router 		/manager/book [POST]
func AddBook(c *gin.Context) {

}

// DeleteBook godoc
//
// @Summary		删除图书信息
// @Description 管理员删除指定书名的图书信息
// @Tags manager
// @Router 		/manager/book/{book_name} [DELETE]
func DeleteBook(c *gin.Context) {

}
