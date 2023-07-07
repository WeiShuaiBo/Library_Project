package logic

import (
	"Library-management/appv1/modle"
	"Library-management/appv1/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Create godoc
// @Summary 添加书记
// @Description 管理员可以输入图书的所有信息来增加图书
// @Accept mpfd
// @Produce json
// @Tags administrator
// @Param book_name formData string true "Name of the book"
// @Param author formData string true "Author of the book"
// @Param publishing_house formData string true "Publishing house of the book"
// @Param translator formData string true "Translator of the book"
// @Param publish_date formData string true "Publish date of the book (format: 2006-01-02)"
// @Param pages formData string true "Number of pages in the book"
// @Param IBSN formData string true "IBSN of the book"
// @Param price formData string true "Price of the book"
// @Param brief_introduction formData string true "Brief introduction of the book"
// @Param author_introduction formData string true "Author introduction of the book"
// @Param img_url formData string true "URL of the book cover image"
// @Param del_flg formData string true "Deletion flag of the book (0 or 1)"
// @Router /all/ad/book [post]
func Create(c *gin.Context) {
	var book modle.Book
	book_name := c.PostForm("book_name")
	author := c.PostForm("author")
	publishing_house := c.PostForm("publishing_house")
	translator := c.PostForm("translator")
	publish_date := c.PostForm("publish_date")
	pages, _ := strconv.ParseInt(c.PostForm("pages"), 10, 64)
	IBSN := c.PostForm("IBSN")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	brief_introduction := c.PostForm("brief_introduction")
	author_introduction := c.PostForm("author_introduction")
	img_url := c.PostForm("img_url")
	del_flg, _ := strconv.ParseInt(c.PostForm("del_flg"), 10, 64)
	book.BookName = book_name
	book.Author = author
	book.PublishingHouse = publishing_house
	book.Translator = translator
	layout := "2006-01-02"
	book.PublishDate, _ = time.Parse(layout, publish_date)
	book.Pages = int(pages)
	book.ISBN = IBSN
	book.Price = price
	book.BriefIntroduction = brief_introduction
	book.AuthorIntroduction = author_introduction
	book.ImgURL = img_url
	book.DelFlg = int(del_flg)
	if err := modle.GlobalConn.Create(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, tools.HttpCode{
			Message: "图书添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Message: "图书添加成功",
	})
}

// Delete godoc
// @Summary 删除书籍
// @Description 管理员可以根据书籍名字来删除书
// @tags administrator
// @Accept json
// @Produce json
// @Param book_name path string true "Name of the book to delete"
// @Router /all/ad/book/{book_name} [delete]
func Delete(c *gin.Context) {
	var books modle.Book
	book_name := c.Param("book_name")
	// 从数据库中获取指定名称的图书信息
	if err := modle.GlobalConn.Where("book_name = ?", book_name).First(&books).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No book found",
		})
		return
	}
	if err := modle.GlobalConn.Where("book_name = ?", book_name).Delete(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "删除书籍失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "删除书记成功",
	})
}

// GetBookUser godoc
// @Summary 查询书的借阅信息
// @Description  管理员可以根据书籍的名字来查看书的借阅信息
// @Tags administrator
// @Accept json
// @Produce json
// @Param book_name path string true "Book Name"
// @Router /all/ad/book/{book_name} [get]
func GetBookUser(c *gin.Context) {
	var books modle.Book
	book_name := c.Param("book_name")
	// 从数据库中获取指定名称的图书信息
	if err := modle.GlobalConn.Where("book_name= ?", book_name).First(&books).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No book found",
		})
		return
	}
	var borrows modle.Borrows
	borrows.BookId = books.BookID
	if err := modle.GlobalConn.Where("book_id=?", borrows.BookId).First(&borrows).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "借阅信息未查到",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": borrows,
	})

}

// GetUserBook godoc
//
// @Summary  借阅信息
// @Description 管理员可以根据用户的名子来查看用户的借阅信息
// @Tags administrator
// @Accept json
// @Produce json
// @Param name path string true "Username"
// @Router /all/ad/user/{name} [get]
func GetUserBook(c *gin.Context) {
	var users modle.User
	name := c.Param("name")
	fmt.Printf("name:%s\n", name)
	// 从数据库中获取指定名称的图书信息
	if err := modle.GlobalConn.Where("name = ?", name).First(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No user found",
		})
		return
	}
	var borrows modle.Borrows
	borrows.UserId = users.UserId
	if err := modle.GlobalConn.Where("user_id=?", borrows.UserId).First(&borrows).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "借阅信息未查到",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": borrows,
	})

}

// GetUserBooks godoc
// @Summary 查看所有借阅信息
// @Description 管理园可以查看所有的借阅信息
// @Tags administrator
// @Produce json
// @Router /all/ad/users [get]
func GetUserBooks(c *gin.Context) {
	var borrows []modle.Borrows
	//从数据库中获取所有图书的信息
	if err := modle.GlobalConn.Find(&borrows).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	// 将图书信息返回给客户端
	c.JSON(http.StatusOK, gin.H{"data": borrows})
}
