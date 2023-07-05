package logic

import (
	"Library/appV0/model"
	"Library/appV0/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreatLibrary godoc
//
//	@Summary		添加图书
//	@Description	会执行添加图书操作
//	@Tags			user/admin
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			title	formData	string	true	"书名"
//	@Param			author		formData	string	true	"作者"
//	@Param			publisher		formData	string	true	"出版社"
//	@Param			edition		formData	string	true	"版号"
//	@Param			stock		formData	int	true	"库存"
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/user/admin/creatLibrary [POST]
func CreatLibrary(c *gin.Context) {
	userId := int64(0)
	ui, exists := c.Get("userId")
	if exists {
		userId = ui.(int64)
	}

	fmt.Println("userId:", userId)

	privilegeById := model.FindUserPrivilegeById(userId)

	if privilegeById == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserPrivilege,
			Message: "权限不够",
			Data:    struct{}{},
		})
		return
	}
	ret := model.Library{}
	err := c.ShouldBind(&ret)
	fmt.Println(ret)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "图书信息错误",
			Data:    struct{}{},
		})
		return
	}

	library := &model.Library{
		Title:     ret.Title,
		Author:    ret.Author,
		Publisher: ret.Publisher,
		Edition:   ret.Edition,
		Stock:     ret.Stock,
		UserId:    userId,
	}

	fmt.Println(library)

	err = model.CreatLibrary(library)

	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "添加数据库失败",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.UserInfoErr,
		Message: "添加成功",
	})

}

// GetLibrary godoc
//
//	@Summary		查询所有图书
//	@Description	会执行查询全部图书操作
//	@Tags			guest
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			page	formData	int	true	"页码"
//	@Param			pageSize		formData	int	true	"页码大小"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/guest/getLibrary [POST]
func GetLibrary(c *gin.Context) {
	page := model.Page{}
	err := c.ShouldBind(&page)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "页码解析错误",
			Data:    struct{}{},
		})
		return
	}

	library, _ := model.GetLibrary(page.Page, page.PageSize)

	fmt.Println("library:", library)

	c.JSON(http.StatusOK, tools.HttpCode{
		Code: tools.OK,
		Data: library,
	})

}

// UpdateLibrary godoc
//
//	@Summary		修改图书
//	@Description	会执行修改图书操作
//	@Tags			user/admin
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id	formData	int64	true	"id"
//	@Param			title	formData	string	true	"书名"
//	@Param			author		formData	string	true	"作者"
//	@Param			publisher		formData	string	true	"出版社"
//	@Param			edition		formData	string	true	"版号"
//	@Param			stock		formData	int	true	"库存"
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/user/admin/updateLibrary [PUT]
func UpdateLibrary(c *gin.Context) {
	userId := int64(0)
	ui, exists := c.Get("userId")
	if exists {
		userId = ui.(int64)
	}

	fmt.Println("userId:", userId)

	privilegeById := model.FindUserPrivilegeById(userId)
	if privilegeById == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserPrivilege,
			Message: "权限不够",
			Data:    struct{}{},
		})
		return
	}

	ret := model.Library{}
	err := c.ShouldBind(&ret)
	fmt.Println(ret)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "图书信息错误",
			Data:    struct{}{},
		})
		return
	}

	library := &model.Library{
		Id:        ret.Id,
		Title:     ret.Title,
		Author:    ret.Author,
		Publisher: ret.Publisher,
		Edition:   ret.Edition,
		Stock:     ret.Stock,
		UserId:    userId,
	}

	fmt.Println(library)

	err = model.UpdateLibrary(library)

	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "添加数据库失败",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.UserInfoErr,
		Message: "修改成功",
	})

}

// DeleteLibrary godoc
//
//	@Summary		删除图书
//	@Description	根据ID删除图书
//	@Tags			user/admin
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			library_id	query	string	true	"id"
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,404,500	{object}	tools.HttpCode{data=Token}
//	@Router			/user/admin/deleteLibrary [DELETE]
func DeleteLibrary(c *gin.Context) {
	userId := int64(0)
	ui, exists := c.Get("userId")
	if exists {
		userId = ui.(int64)
	}

	fmt.Println("userId:", userId)

	privilegeById := model.FindUserPrivilegeById(userId)

	if privilegeById == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserPrivilege,
			Message: "权限不够",
			Data:    struct{}{},
		})
		return
	}

	libraryIdStr := c.Query("library_id")

	if libraryIdStr == "" {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "获取图书Id为空",
			Data:    libraryIdStr,
		})
		return
	}

	// 将字符串转换为int64
	libraryId, err := strconv.ParseInt(libraryIdStr, 10, 64)

	if err != nil {
		fmt.Println("无法转换libraryId为int64:", err)
		return
	}

	// 打印结果
	fmt.Println("libraryId:", libraryId)

	err1 := model.DeleteLibrary(libraryId)

	if err1 != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "删除失败",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "删除成功",
		Data:    struct{}{},
	})
}

// FindLibrary godoc
//
//	@Summary		查找图书
//	@Description	根据title查找图书
//	@Tags			guest
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			title	query	string	true	"书名"
//	@response		200,404,500	{object}	tools.HttpCode{data=model.Library}
//	@Router			/guest/findLibrary [POST]
func FindLibrary(c *gin.Context) {

	title := c.Query("title")
	fmt.Println("title:", title)
	library, err1 := model.FindLibrary(title)

	if err1 != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "未查到",
			Data:    struct{}{},
		})
		return
	}

	fmt.Println("library:", library)

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "",
		Data:    library,
	})
}
