package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/appV5/model"
	"library/appV5/tools"
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
//	@Param			pageNumber			path		int	false	"int valid"	minimum(1)
//
// @Success 200 {object} tools.HttpCode
// @Failure 404 {object} tools.HttpCode
//
//	@Router			/{pageNumber} [get]
func GetAll(c *gin.Context) {
	pageNumber := c.Param("pageNumber")
	fmt.Println(pageNumber)
	pageNumberInt, _ := strconv.Atoi(pageNumber)
	fmt.Printf("-----")
	//pageNumberInt, _ := strconv.ParseInt(pageNumber, 10, 64)
	fmt.Printf("%d", pageNumberInt)
	fmt.Printf("+++++++++++++++++")
	ret := model.GetAllRedis(pageNumberInt)

	if len(ret) == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "查询页码过大",
			Data:    ret,
		})
		return
	}
	fmt.Printf("222222222222222222")
	if ret[0].Id > 0 {
		//fmt.Printf("book:%+v\n", ret)
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "查询成功",
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

// GetBookPhotoByName godoc
//
//	@Summary		根据id获得图书
//	@Description	根据ID获得图书
//	@Tags			tourist
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			bookName			path		int	true	"int valid"	minimum(1)
//	@response		200,404,500	{object}	tools.HttpCode{data=model.Book}
//	@Router			/book/{bookName} [get]
func GetBookPhotoByName(c *gin.Context) {
	fmt.Printf("GetBookPhotoByName进入成功")
	fmt.Printf("+++++++++++++++++++++++++++++++++++++++++++")
	// 获取bookName参数
	bookName := c.Param("bookName")
	fmt.Printf(bookName)
	// 判断是否获取到bookName
	if bookName == "" {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "获取bookName失败",
			Data:    struct{}{},
		})
		return
	}
	// 查询数据库
	fmt.Printf("准备前往redis")
	ret := model.GetBookByNameRedis(bookName)
	//ret := model.GetBookByName(bookName)
	// 判断是否查询到数据
	fmt.Print(ret)
	if ret != nil {
		fmt.Printf("book:%+v\n", ret)
		//dir, _ := os.Getwd()
		//fmt.Printf(dir)
		filePath := "./appV0/static/images/" + ret.ImgUrl // 图片文件路径
		c.File(filePath)
		return
	}

	c.JSON(http.StatusNotFound, tools.HttpCode{
		Code:    tools.NotFound,
		Message: "查询数据库失败",
		Data:    struct{}{},
	})
	return
}

// Qrcode godoc
//
//	@Summary		获取二维码
//	@Description	获取二维码
//	@Tags			tourist
//	@Accept			json
//	@Produce		json
//
// @Success 200 {object} tools.HttpCode
// @Failure 404 {object} tools.HttpCode
//
//	@Router			/qrcode [get]
func Qrcode(c *gin.Context) {
	code := generateCode()
	url := "./appV2/static/qrcode/" + code + ".png"
	fmt.Print(url)
	fmt.Print(code)
	if model.GetQRcode(code, url) {
		c.File(url)
	}

	if model.QrcodeRedis(code) {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "二维码生成成功",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusNotFound, tools.HttpCode{
		Code:    tools.NotFound,
		Message: "二维码生成失败",
		Data:    struct{}{},
	})
	return

}

// AutomaticLogin godoc
//
//	@Summary		自动登录
//	@Description	会执行自动登录操作
//	@Tags			tourist
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			code			path		int	true	"int valid"	minimum(1)
//
// @response		200,500	{object}	tools.HttpCode{data=Token}
// @Router			/AutomaticLogin/{code} [Get]
func AutomaticLogin(c *gin.Context) {
	code := c.Param("code")
	ret := model.AutomaticLogin(code)
	if ret == nil {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "获取失败",
			Data:    struct{}{},
		})
	}
	//	下发Token并判断
	a, r, err := tools.Token.GetToken(ret.Id, ret.Name, "user")
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "生成Token失败",
			Data:    struct{}{},
		})
		return
	}
	//	通知成功
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "登录成功，正在跳转",
		Data: Token{
			AccessToken:  a,
			RefreshToken: r,
		},
	})
	return

}
