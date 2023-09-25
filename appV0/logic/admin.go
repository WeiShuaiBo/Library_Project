package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/appV0/model"
	"library/appV0/tools"
	"net/http"
	"strconv"
)

type Admin struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
}
type Book struct {
	BN          string `gorm:"type:varchar(20)" json:"bn"`
	Name        string `gorm:"type:varchar(200)" json:"name"`
	Description string `gorm:"type:varchar(15000)"`
	Count       int    `json:"count"`
	ImgUrl      string `json:"img_url" gorm:"varchar(200)"`
}

// LibrarianLogin godoc
//
//	@Summary		管理园登录
//	@Description	会执行管理员登录操作
//	@Tags			admin
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			pwd		formData	string	true	"密码"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/adminLogin [POST]
func LibrarianLogin(c *gin.Context) {
	var admin Admin
	//	绑定并判断
	if c.ShouldBind(&admin) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "绑定失败",
			Data:    struct{}{},
		})
		return
	}
	//	查询数据库并判断
	dbUser := model.GetAdmin(admin.Name, admin.Pwd)
	if dbUser.Id > 0 {
		err := model.SetSession(c, dbUser.Name, dbUser.Id)
		c.SetCookie("name", dbUser.Name, 3600, "/", "", false, true) //domain 写域名的话 会导致IP访问无效
		c.SetCookie("id", strconv.FormatInt(dbUser.Id, 10), 3600, "/", "", false, true)
		if err != nil {
			c.JSON(http.StatusOK, tools.HttpCode{
				Code:    tools.UserInfoErr,
				Message: err.Error(),
			})
		}

		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "登录成功，整在跳转~",
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.UserInfoErr,
		Message: "用户信息错误",
	})
}

// AdminLogout godoc
//
//	@Summary		管理员退出
//	@Description	会执行管理员退出操作
//	@Tags			admin
//	@Accept			multipart/form-data
//	@Produce		json
//	@response		500,401	{object}	tools.HttpCode
//	@Router			/adminLogout [get]
func AdminLogout(c *gin.Context) {
	//设置登录态
	_ = model.FlushSession(c)
	c.SetCookie("name", "", 3600, "/", "", false, true) //domain 写域名的话 会导致IP访问无效
	c.SetCookie("id", "", 3600, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/adminLogin")
	return

	c.JSON(http.StatusUnauthorized, tools.HttpCode{
		Code: tools.OK,
		Data: struct{}{},
	})
	return
}

// GetRecords godoc
//
//	@Summary		获取借书记录列表
//	@Description	获取借书记录列表列表
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@response		200,500	{object}	tools.HttpCode{data=[]model.Record}
//	@Router			/admin/records/getRecords [get]
func GetRecords(c *gin.Context) {
	fmt.Printf("进入logic_get")
	ret := model.GetRecords()
	if ret == nil {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "查询数据库有问题",
			Data:    struct{}{},
		})
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code: tools.OK,
		Data: ret,
	})
}

// GetUserRecordStatus godox
// @Summary 查看用户借书记录
// @Description 根据用户ID获取用户借书信息
// @Tags admin
// @Accept multipart/form-data
// @Produce json
// @Param userId path int true "用户ID"
// @Success 200 {object} tools.HttpCode
// @Router /admin/records/status/{id} [get]
func GetUserRecordStatus(c *gin.Context) {
	statusIdStr := c.Param("id")
	statusId, _ := strconv.ParseInt(statusIdStr, 10, 64)
	if statusId <= 0 {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "id获取失败",
			Data:    struct{}{},
		})
		return
	}
	ret := model.GetUserRecordStatus(statusId)
	if ret[0].Id > 0 {
		fmt.Printf("UserRecord:%+v\n", ret)
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "",
			Data:    ret,
		})
		return
	}

	c.JSON(http.StatusNotFound, tools.HttpCode{
		Code: tools.NotFound,
		Data: struct{}{},
	})
	return
}

// GetUserInformation godoc
// @Summary 查看个人信息
// @Description 根据用户ID获取用户信息
// @Tags admin
// @Accept multipart/form-data
// @Produce json
// @Param userId path int true "用户ID"
// @Success 200 {object} tools.HttpCode
// @Router /admin/records/GetPersonalInformation/{userId} [get]
func GetUserInformation(c *gin.Context) {
	statusIdStr := c.Param("userId")
	statusId, _ := strconv.ParseInt(statusIdStr, 10, 64)
	if statusId <= 0 {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "id获取失败",
			Data:    struct{}{},
		})
		return
	}
	ret := model.GetUserInformation(statusId)
	if ret.Id > 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "",
			Data:    ret,
		})
		return
	}

	c.JSON(http.StatusNotFound, tools.HttpCode{
		Code: tools.NotFound,
		Data: struct{}{},
	})
	return
}

// AddBook godoc
//
//	@Summary		增添图书
//	@Description	会执行增添图书操作
//	@Tags			admin
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			BN	formData	string	true	"种类"
//	@Param			Name		formData	string	true	"书名"
//	@Param			Description		formData	string	true	"介绍"
//	@Param			Count		formData	int	true	"数量"
//	@Param			ImgUrl		formData	string	true	"imgUrl"
//	@response		200,500	{object}	tools.HttpCode
//	@Router			/admin/books/AddBook [POST]
func AddBook(c *gin.Context) {
	book := &Book{}
	//	绑定并判断
	if c.ShouldBind(&book) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "绑定失败",
			Data:    struct{}{},
		})
		return
	}
	ret := model.AddBook(book.BN, book.Name, book.Description, book.ImgUrl, book.Count)
	if ret == 1 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "查询数据库有问题",
			Data:    struct{}{},
		})
		return
	} else {
		if ret == 2 {
			c.JSON(http.StatusOK, tools.HttpCode{
				Code:    tools.OK,
				Message: "用户名已存在",
				Data:    struct{}{},
			})
		}
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "注册成功",
		Data:    struct{}{},
	})
}

// DeleteBook godoc
// @Summary 删除图书
// @Description 根据bookID删除图书
// @Tags admin
// @Accept multipart/form-data
// @Produce json
// @Param bookId path int true "bookID"
// @Success 200 {object} tools.HttpCode
//
//	@Router			/admin/books/DeleteBook/{bookId} [get]
func DeleteBook(c *gin.Context) {
	statusIdStr := c.Param("bookId")
	statusId, err := strconv.ParseInt(statusIdStr, 10, 64)
	if err != nil {
		fmt.Println("解析失败:", err)
	} else {
		fmt.Println(statusId)
	}
	if statusId <= 0 {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "id获取失败",
			Data:    struct{}{},
		})
		return
	}
	ret := model.DeleteBook(statusId)
	if ret == 1 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "查询数据库有问题",
			Data:    struct{}{},
		})
		return
	} else {
		if ret == 2 {
			c.JSON(http.StatusOK, tools.HttpCode{
				Code:    tools.OK,
				Message: "图书不存在",
				Data:    struct{}{},
			})
		}
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "删除图书成功",
		Data:    struct{}{},
	})
}

// GETBookRecord godoc
// @Summary 查看图书借还记录
// @Description 根据bookID查看图书借还记录
// @Tags admin
// @Accept multipart/form-data
// @Produce json
// @Param bookId path int true "bookID"
// @Success 200 {object} tools.HttpCode
//
//	@Router			/admin/records/GETBookRecord/{bookId} [get]
func GETBookRecord(c *gin.Context) {
	statusIdStr := c.Param("bookId")
	statusId, err := strconv.ParseInt(statusIdStr, 10, 64)
	if err != nil {
		fmt.Println("解析失败:", err)
	} else {
		fmt.Println(statusId)
	}
	if statusId <= 0 {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "id获取失败",
			Data:    struct{}{},
		})
		return
	}
	ret := model.GETBookRecord(statusId)
	if ret.Id <= 0 {
		c.JSON(http.StatusNotFound, tools.HttpCode{
			Code:    tools.NotFound,
			Message: "获取失败",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "获取图书借还信息成功",
		Data:    ret,
	})

}
