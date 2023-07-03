package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/appV0/model"
	"library/appV0/tools"
	"net/http"
	"strconv"
)

// GetAll godoc
//
//	@Summary		获取前100书的信息
//	@Description	从数据库获取前100书的信息
//	@Tags			tourist
//	@Accept			json
//	@Produce		json
//
// @Success 200 {object} tools.HttpCode
// @Failure 404 {object} tools.HttpCode
//
//	@Router			/ [get]
func GetAll(c *gin.Context) {
	fmt.Printf("可以走到这里吗？")
	ret := model.GetAll()
	if ret[0].Id > 0 {
		fmt.Printf("book:%+v\n", ret)
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "",
			Data:    ret,
		})
		return
	}
	c.JSON(http.StatusNotFound, tools.HttpCode{
		Code:    tools.NotFound,
		Message: "数据库查询失败",
		Data:    struct{}{},
	})
}
func Html(c *gin.Context) {
	//发送一个带有两个参数的页面给前端
	c.HTML(200, "index.html", gin.H{
		"usernameer'r":  "欢迎登录",
		"usernameError": "ok",
	})

}

// GetBook godoc
//
//	@Summary		根据id获得图书
//	@Description	根据ID获得图书
//	@Tags			tourist
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id			path		int	true	"int valid"	minimum(1)
//	@response		200,404,500	{object}	tools.HttpCode{data=model.Book}
//	@Router			/books/{id} [get]
func GetBook(c *gin.Context) {
	//	获取idStr
	idStr := c.Param("id")
	fmt.Printf(idStr)
	//	转int
	id, _ := strconv.ParseInt(idStr, 10, 64)
	fmt.Printf(string(id))
	//	判断失败？
	if id <= 0 {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "获取id失败",
			Data:    struct{}{},
		})
		return
	}
	//	查数据库
	ret := model.GetBook(id)
	//	根据查出的数据库判断id是否错误
	if ret != nil {
		fmt.Printf("book:%+v\n", ret)
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "查询数据库成功",
			Data:    ret,
		})
		return
	}
	c.JSON(http.StatusNotFound, tools.HttpCode{
		Code:    tools.NotFound,
		Message: "查询数据库失败",
		Data:    struct{}{},
	})
	return
}
