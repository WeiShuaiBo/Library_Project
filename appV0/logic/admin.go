package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/appV0/logger"
	model2 "library/appV0/model"
	"library/appV0/tools"
	"net/http"
	"strconv"
)

// AdminGetBooks
//
//	@Summary		获取列表书籍
//	@Description	会执行获得所有书籍的详细信息
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Response		200,500	{object}	tools.HttpCode{data=[]model.Book}
//	@Router			/admin/books [get]
func AdminGetBooks(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	ret := make([]model2.BookInfo, 0)
	if err := model2.AdminGetBooks(&ret, limit, offset); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.BookErr,
			Message: "查询图书信息失败",
			Data: struct {
			}{},
		})
		logger.Log.Error("查询图书信息失败", err)
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "查询图书信息成功",
		Data:    ret,
	})
	logger.Log.Info("查询图书信息成功")
	return
}

// AdminGetBookByKeyWord
//
//	@Summary		模糊查询书籍
//	@Description	管理员模糊查询书籍的详细信息
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			keyWord	query		string	true	"关键词"
//	@Response		200,500	{object}	tools.HttpCode{data=[]model.BookInfo}
//	@Router			/admin/bookByKeyWord [post]
func AdminGetBookByKeyWord(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	keyWord := c.Query("keyWord")
	ret := make([]model2.BookInfo, 0)
	if err := model2.AdminGetBookByKeyWord(&ret, keyWord, limit, offset); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.BookErr,
			Message: "根据关键词查询图书信息失败",
			Data: struct {
			}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "根据关键词查询图书信息成功",
		Data:    ret,
	})
	return
}

// AdminGetBooksById 根据Id查询书籍
//
//	@Summary		管理员根据Id查询书籍
//	@Description	管理员根据Id查询书籍的详细信息
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"搜索id"
//	@Response		200,500	{object}	tools.HttpCode{data=model.Book}
//	@Router			/admin/book/{id} [get]
func AdminGetBooksById(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	logger.Log.Info("id:", id)

	if id <= 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.AdminInfoErr,
			Message: "id数据有误",
			Data: struct {
			}{},
		})
		logger.Log.Error("id数据有误")
		return
	}

	ret := &model2.BookInfo{}
	if err := model2.AdminGetBooksById(ret, idStr, limit, offset); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.BookErr,
			Message: "根据id查询图书信息失败",
			Data: struct {
			}{},
		})
		logger.Log.Error(err)

		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "根据id查询图书信息成功",
		Data:    ret,
	})
}

// CreatBook
//
//	@Summary		新建图书
//	@Description	管理员新增图书
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			book_name	body		string	true	"图书名字"
//	@Param			author		body		string	true	"图书作者"
//	@Param			isbn		body		string	true	"ISBN编号"
//	@Response		200,400,500	{object}	tools.HttpCode
//	@Router			/admin/book [post]
func CreatBook(c *gin.Context) {
	book := &model2.BookInfo{}
	if c.ShouldBind(&book) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.AdminInfoErr,
			Message: "数据格式有误",
			Data:    nil,
		})
		logger.Log.Error("数据格式有误")
		return
	}
	if err := model2.CreatBook(book); err != nil {
		fmt.Printf("err:%s\n", err)
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.AdminInfoErr,
			Message: "添加图书失败",
			Data: struct {
			}{},
		})
		logger.Log.Error("添加图书失败", err)
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "添加图书成功",
		Data:    book,
	})
	logger.Log.Info("添加图书成功")
	return
}

// UpdateBook
//
//	@Summary		修改图书
//	@Description	管理员修改图书
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int		true	"更新图书id"
//	@Param			book_name		body		string	true	"图书名字"
//	@Param			author		body		string	true	"图书作者"
//	@Param			isbn		body		string	true	"图书编号"
//	@Response		200,400,500	{object}	tools.HttpCode
//	@Router			/admin/book/{id} [put]
func UpdateBook(c *gin.Context) {
	isbnStr := c.Param("isbn")
	book := &model2.BookInfo{}
	if c.ShouldBind(&book) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.AdminInfoErr,
			Message: "数据格式有误",
			Data:    nil,
		})
		logger.Log.Error("数据格式有误")
		return
	}
	if err := model2.UpdateBook(book, isbnStr); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.AdminInfoErr,
			Message: "更新图书失败",
			Data: struct {
			}{},
		})
		logger.Log.Error("更新图书失败", err)
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "更新图书成功",
		Data: struct {
		}{},
	})
	return
}

// DeleteBook
//
//	@Summary		删除书籍
//	@Description	管理员删除书籍
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			isbn			path		int	true	"删除isbn"
//	@Response		200,400,500	{object}	tools.HttpCode
//	@Router			/admin/book/{isbn} [delete]
func DeleteBook(c *gin.Context) {
	isbnStr := c.Param("isbn")
	logger.Log.Info(isbnStr)
	if err := model2.DeleteBook(isbnStr); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.AdminInfoErr,
			Message: "删除图书失败",
			Data: struct {
			}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "删除图书成功",
		Data: struct {
		}{},
	})
	logger.Log.Info("删除图书成功")
	return
}

// AdminGetInfo
//
//	@Summary		借阅记录
//	@Description	会执行获得所有书籍的借阅记录
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Response		200,500	{object}	tools.HttpCode{data=[]model.LendBooks}
//	@Router			/admin/getInfo [get]
func AdminGetInfo(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	ret := &[]model2.UserBook{}
	if err := model2.AdminGetInfo(ret, limit, offset); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "获取借阅表数据失败",
			Data: struct {
			}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "获取所有借阅表数据成功",
		Data:    ret,
	})
	return

}
