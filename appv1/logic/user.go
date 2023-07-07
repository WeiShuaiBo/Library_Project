package logic

import (
	"Library-management/appv1/modle"
	"Library-management/appv1/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// GetUser		godoc
// @Summary 		用户查阅信息
// @Description 	    用户登录过后可以查看自己的信息
// @Tags 			user
// @Param debug  header string false "调试模式" default(123)
// @Router		/all/user [post]
func GetUser(c *gin.Context) {
	//userId := modle.GetCurrentUserID(c)
	//userId, _ := c.Get("userId")
	userId := 33
	user, err := modle.GetUserInfo(userId)
	if err != nil {
		c.JSON(404, tools.HttpCode{
			Message: "出错了",
		})
		return
	}
	c.JSON(200, tools.HttpCode{
		Data: user})
}

// UpdateUser godoc
// @Summary  更新用户信息
// @Description  用户登录过后可以修改自己的用户信息
// @Tags user
// @Accept			mpfd
// @Produce		  json
// @Param debug  header string false "调试模式" default(123)
// @Param name formData string true "The name of the user"
// @Param password formData string true "The password for the user account"
// @Param phone formData string true "The phone number of the user"
// @Router /all/user [put]
func UpdateUser(c *gin.Context) {
	user := &modle.User{}
	user.UserId = modle.GetCurrentUserID(c) //测试用的
	//userId, _ := c.Get("userId")
	if err := c.ShouldBind(&user); err != nil || user.UserId <= 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Message: "数据解析失败,Id不能为空",
			Data:    struct{}{},
		})
		return
	}
	fmt.Printf("user:%+v\n", user)
	err := modle.UpdataUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "更新失败"})
		return
	}
	c.JSON(200, gin.H{"Message": "用户更新成功"})
}

// Borrow godoc
// @Summary 借书
// @Description 用户进行借书
// @Tags user
// @Accept json
// @Produce json
// @Param book_name path string true "Book Name"
// @Router /all/borrow/{book_name} [post]
func Borrow(c *gin.Context) {
	var book modle.Book
	tx := modle.GlobalConn.Begin()
	book.BookName = c.Param("book_name")

	// 根据指定的标题从数据库中获取图书信息
	if err := tx.Where("book_name = ?", book.BookName).Take(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "未找到该书",
		})
		tx.Rollback()
		return
	}

	userID := modle.GetCurrentUserID(c) //这里是用来测试的
	//userID:=c.Get("userId")
	fmt.Printf("User ID: %s, Book ID: %s\n", userID, book.BookID)

	borrow := &modle.Borrows{
		UserId:     userID,
		BookId:     book.BookID,
		BorrowDate: time.Now(),
		ReturnDate: time.Now().AddDate(0, 1, 0),
	}

	if err := tx.Create(&borrow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "借书失败"})
		tx.Rollback()
		return
	}

	if book.DelFlg == 0 {
		book.DelFlg = 1
		if err := tx.Model(&modle.Book{}).Where("book_id = ?", book.BookID).Updates(book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "更新书籍数量失败"})
			tx.Rollback()
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "该书本已经没有了"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "借书成功"})
	tx.Commit()
}

// Return godoc
// @Summary 还书
// @Description 根据图书名称归还图书
// @Tags user
// @Accept json
// @Produce json
// @Param book_name path string true "图书名称"
// @Router /all/return/{book_name} [post]
func Return(c *gin.Context) {
	var book modle.Book
	book.BookName = c.Param("book_name")
	tx := modle.GlobalConn.Begin()

	// 根据指定的标题从数据库中获取图书信息
	if err := tx.Where("book_name = ?", book.BookName).Take(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "未找到该书",
		})
		tx.Rollback()
		return
	}
	var borrows modle.Borrows
	borrows.BookId = book.BookID
	if err := modle.GlobalConn.Where("book_id = ?", borrows.BookId).First(&borrows).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "没有查到该用户的借书记录",
		})
		tx.Rollback()
		return
	}
	borrows.ReturnDate = time.Now()
	modle.GlobalConn.Model(&modle.Borrows{}).Where("borrow_id = ?", borrows.BorrowId).Updates(borrows)
	book.DelFlg = 0
	modle.GlobalConn.Model(&modle.Book{}).Where("book_id = ?", book.BookID).Updates(book)
	c.JSON(http.StatusOK, gin.H{"Message": "还书成功"})
	tx.Commit()
}

// Logout godoc
//
//	@Summary		用户退出
//	@Description	会执行用户退出操作
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Router			/all/logout [get]
func Logout(c *gin.Context) {
	_ = modle.FlushSession(c)
	//这里暂时先不改为 401，有些接口确实不需要登录态
	c.JSON(http.StatusUnauthorized, tools.HttpCode{
		Code: tools.OK,
		Data: struct{}{},
	})
	return
	c.JSON(http.StatusOK, tools.HttpCode{
		Message: "退出成功",
	})
}

// Borrowing godoc
// @Summary 借阅记录
// @Description 用户可以查看自己的借阅记录
// @Tags user
// @Produce json
// @Router /all/borrowing [get]
func Borrowing(c *gin.Context) {
	userId := modle.GetCurrentUserID(c)
	//userId := c.GetInt("userId")
	var borrows modle.Borrows

	borrows.UserId = userId
	// 从数据库中获取指定名称的图书信息
	if err := modle.GlobalConn.Where("user_id = ?", borrows.UserId).First(&borrows).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "未查到该用户的借阅信息",
		})
		return
	}
	// 将图书信息返回给客户端
	c.JSON(http.StatusOK, gin.H{"data": borrows})
}
