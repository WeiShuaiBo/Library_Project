package logic

import (
	"Library/appV0/model"
	"Library/appV0/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetRecordByUserId godoc
//
//	@Summary		借书记录
//	@Description	根据用户id查找借书记录
//	@Tags			user/common
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,404,500	{object}	tools.HttpCode{data=model.Library}
//	@Router			/user/common/getRecordByUserId [GET]
func GetRecordByUserId(c *gin.Context) {

	userId := int64(0)
	ui, exists := c.Get("userId")
	if exists {
		userId = ui.(int64)
	}

	record, err := model.GetRecordByUserId(userId)

	fmt.Println("record:", record)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "未查到",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "",
		Data:    record,
	})

}

// GetARecordByUserId godoc
//
//	@Summary		借书记录(单个用户)
//	@Description	根据用户id查找借书记录
//	@Tags			user/admin
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			user_id	query	int64	true	"用户Id"
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,404,500	{object}	tools.HttpCode{data=model.Library}
//	@Router			/user/admin/getARecordByUserId [POST]
func GetARecordByUserId(c *gin.Context) {
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

	userIdStr := c.Query("user_id")
	ud, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		fmt.Println("转换失败")
		return
	}

	record, err1 := model.GetRecordByUserId(ud)

	if err1 != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "未查到",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "",
		Data:    record,
	})
}

// GetRecord godoc
//
//	@Summary		借书记录（全部）
//	@Description	查找全部借书记录
//	@Tags			user/admin
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,404,500	{object}	tools.HttpCode{data=model.Library}
//	@Router			/user/admin/getRecord [GET]
func GetRecord(c *gin.Context) {

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

	records, err := model.GetRecord()
	fmt.Println("records:", records)
	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "未查到",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "",
		Data:    records,
	})
}

// GetExpected godoc
//
//	@Summary		预期用户（全部）
//	@Description	查找全部预期
//	@Tags			user/admin
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,404,500	{object}	tools.HttpCode{data=model.Library}
//	@Router			/user/admin/getExpected [GET]
func GetExpected(c *gin.Context) {

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

	borrows, err := model.GetExpected()

	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "未查到",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "",
		Data:    borrows,
	})
}
