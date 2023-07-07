package logic

import (
	"Library-management/appv1/modle"
	"Library-management/appv1/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Book struct {
	BookID             uint      `json:"book_id"`
	BookName           string    `json:"book_name"`
	Author             string    `json:"author"`
	PublishingHouse    string    `json:"publishing_house"`
	Translator         string    `json:"translator"`
	PublishDate        time.Time `json:"publish_date"`
	Pages              int       `json:"pages"`
	ISBN               string    `json:"isbn"`
	Price              float64   `json:"price"`
	BriefIntroduction  string    `json:"brief_introduction"`
	AuthorIntroduction string    `json:"author_introduction"`
	ImgURL             string    `json:"img_url"`
	DelFlg             int       `json:"del_flg"`
}

// GetBooks godoc
// @Summary 游客可以查看所有的图书
// @Tags all
// @Router /all/getBooks [GET]
// @Param page query string false "页码" default(1)
// @Produce json
// @Success 200 {object} tools.HttpCode
// @Failure 404 {object} tools.HttpCode
func GetBooks(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.ParseInt(pageStr, 10, 64)

	books, err := modle.GetAllBooks(page)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code: tools.OK,
		Data: books,
	})
}

// GetBook godoc
// @Summary 可以根据某一本图书的名字，作者查询到书的所有信息
// @Tags 	all
// @Accept json
// @Produce json
// @Param 	book_name path string ture "书名"
// @Router	/all/getBook/{book_name} [GET]
func GetBook(c *gin.Context) {
	var books modle.Book
	books.BookName = c.Param("book_name")
	// 从数据库中获取指定名称的图书信息
	books, err := modle.GetABook(books.BookName)
	if err != nil {
		c.JSON(404, tools.HttpCode{
			Message: "查询失败",
		})
		return
	}
	// 将图书信息返回给客户端
	c.JSON(http.StatusOK, tools.HttpCode{
		Message: "查询成功",
		Data:    books,
	})
}
