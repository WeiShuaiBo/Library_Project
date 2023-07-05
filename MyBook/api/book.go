// @Author	zhangjiaozhu 2023/7/4 10:23:00
package api

import (
	"MyBook/common/Response"
	"MyBook/models"
	"MyBook/routers/midwares"
	"MyBook/service"
	"github.com/gin-gonic/gin"
)

type BookReq struct {
	BookName  string `json:"book_name"` // 书名
	Author    string `json:"author"`    // 作者
	Price     string `json:"price"`     // 价格
	Type      string `json:"type"`      // 类型
	Remaining string `json:"remaining"` // 剩余本数
	Desc      string `json:"desc"`      // 简介
}

func GetBook(c *gin.Context) {
	data, err := service.GetBook()
	if err != nil {
		Response.Error(c, "服务器繁忙")
		return
	}
	Response.Success(c, "success", data)
}

// FindBookByName 根据书名查找书籍
func FindBookByName(c *gin.Context) {
	type BookName struct {
		Name string `json:"book_name"`
	}
	var req BookName
	if err := c.ShouldBindJSON(&req); err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	book, err := service.FindBookByName(req.Name)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "success", book)
}

// 添加书籍
func CreateBook(c *gin.Context) {
	var bookReq BookReq
	if err := c.ShouldBindJSON(&bookReq); err != nil {
		Response.Error(c, "请求参数有误")
		return
	}
	// 判断书籍是否存在
	bookResult, err := service.FindBookByName(bookReq.BookName)
	if err != nil {
		Response.Error(c, "服务器繁忙")
		return
	}
	if bookResult.BookId > 0 {
		Response.Error(c, "书籍已存在")
		return
	}
	// 添加书籍
	var book models.Book
	book = models.Book{
		BookName:  bookReq.BookName,
		Author:    bookReq.Author,
		Price:     bookReq.Price,
		Type:      bookReq.Type,
		Remaining: bookReq.Remaining,
		Desc:      bookReq.Desc,
	}
	err = service.CreateBook(&book)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "success", nil)

}

// 删除书籍
func DeleteBookByName(c *gin.Context) {
	type BookReq struct {
		Name string `json:"book_name"`
	}
	var bookReq BookReq
	if err := c.ShouldBindJSON(&bookReq); err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	book, err := service.FindBookByName(bookReq.Name)
	if err != nil {
		Response.Error(c, "服务器繁忙")
		return
	}
	if book.BookId == 0 {
		Response.Error(c, "书籍不存在")
		return
	}
	err = service.DeleteBookByName(book)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "success", nil)

}

// 更新书籍

// 借书
func BorrowBook(c *gin.Context) {
	// 上下文中获取用户id
	userId, exists := c.Get(midwares.ContextUserIDKey)
	if !exists {
		Response.Error(c, "用户未登录")
	}
	type BookReq struct {
		Name string `json:"book_name"`
	}
	var bookReq BookReq
	if err := c.ShouldBindJSON(&bookReq); err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	value, _ := userId.(uint64)
	err := service.BorrowBook(bookReq.Name, value)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "借书成功", nil)
}

// 还书
func ReturnBook(c *gin.Context) {
	// 上下文中获取用户id
	userId, exists := c.Get(midwares.ContextUserIDKey)
	if !exists {
		Response.Error(c, "用户未登录")
	}
	type BookReq struct {
		Name string `json:"book_name"`
	}
	var bookReq BookReq
	if err := c.ShouldBindJSON(&bookReq); err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	value, _ := userId.(uint64)
	err := service.ReturnBook(bookReq.Name, value)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.Success(c, "还书成功", nil)
}

func UpdateBook(c *gin.Context) {
	type Req struct {
		BookName  string // 书名
		Author    string // 作者
		Price     string // 价格
		Type      string // 类型
		Reserve   string // 预约数量
		Loan      string // 借出数量
		Remaining string // 剩余本数
		Desc      string // 简介
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		Response.Error(c, "参数不合法")
		return
	}
	var book *models.Book
	service.UpdateBook(book)
}
