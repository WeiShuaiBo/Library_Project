package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/logger"
	"library/model"
	"library/tools"
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
	ret := make([]model.Book, 0)
	if err := model.AdminGetBooks(&ret); err != nil {
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
//	@Response		200,500	{object}	tools.HttpCode{data=[]model.Book}
//	@Router			/admin/books [get]
func AdminGetBookByKeyWord(c *gin.Context) {
	keyWord := c.Query("keyWord")
	ret := make([]model.Book, 0)
	if err := model.AdminGetBookByKeyWord(&ret, keyWord); err != nil {
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

	ret := &model.Book{}
	if err := model.AdminGetBooksById(ret, id); err != nil {
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
//	@Param			name		body		string	true	"图书名字"
//	@Param			author		body		string	true	"图书作者"
//	@Param			number		body		string	true	"图书编号"
//	@Response		200,400,500	{object}	tools.HttpCode
//	@Router			/admin/book [post]
func CreatBook(c *gin.Context) {
	book := &model.Book{}
	if c.ShouldBind(&book) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.AdminInfoErr,
			Message: "数据格式有误",
			Data:    nil,
		})
		logger.Log.Error("数据格式有误")
		return
	}
	if err := model.CreatBook(book); err != nil {
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
		Data: struct {
		}{},
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
//	@Param			name		body		string	true	"图书名字"
//	@Param			author		body		string	true	"图书作者"
//	@Param			number		body		string	true	"图书编号"
//	@Param			lend_out	body		string	false	"图书作者"
//	@Param			user_id		body		string	false	"图书编号"
//	@Response		200,400,500	{object}	tools.HttpCode
//	@Router			/admin/book/{id} [put]
func UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	book := &model.Book{}
	if c.ShouldBind(&book) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.AdminInfoErr,
			Message: "数据格式有误",
			Data:    nil,
		})
		logger.Log.Error("数据格式有误")
		return
	}
	book.Id = id
	if err := model.UpdateBook(book); err != nil {
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
		Data:    book,
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
//	@Param			id			path		int	true	"删除id"
//	@Response		200,400,500	{object}	tools.HttpCode
//	@Router			/admin/book/{id} [delete]
func DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	logger.Log.Info(id)

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

	book := &model.Book{}
	if err := model.DeleteBook(book, id); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.AdminInfoErr,
			Message: "删除图书失败",
			Data: struct {
			}{},
		})
		logger.Log.Error("删除图书失败", err)
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
	ret := &[]model.LendBooks{}
	if err := model.AdminGetInfo(ret); err != nil {
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
