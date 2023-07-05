package logic

import (
	"github.com/gin-gonic/gin"
	"library/model"
	"library/tools"
	"net/http"
)

// GetBooks godoc
//
//	@Summary		获取图书列表
//	@Description	获取所有图书列表
//	@Tags			Comm
//	@Accept			json
//	@Produce		json
//	@response		200,500	{object}	tools.HttpCode{data=[]model.APIBook}
//	@Router			/books [get]
func GetBooks(c *gin.Context) {
	ret := make([]model.APIBook, 0)
	if err := model.GetBooks(&ret); err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.BookErr,
			Message: "查询图书信息失败",
			Data: struct {
			}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "查询图书信息成功",
		Data:    ret,
	})
}

// GetBookByKeyWord godoc
//
//	@Summary		模糊查询图书
//	@Description	根据关键词模糊查询所有图书
//	@Tags			Comm
//	@Accept			json
//	@Produce		json
//	@Param			keyWord	query		string	true	"关键词"
//	@response		200,500	{object}	tools.HttpCode{data=[]model.APIBook}
//	@Router			/bookByKeyWord [post]
func GetBookByKeyWord(c *gin.Context) {
	keyWord := c.Query("keyWord")
	ret := make([]model.APIBook, 0)
	if err := model.GetBookByKeyWord(&ret, keyWord); err != nil {
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
