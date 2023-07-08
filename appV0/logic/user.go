package logic

import (
	"github.com/gin-gonic/gin"
	"library/appV0/logger"
	model2 "library/appV0/model"
	tools2 "library/appV0/tools"
	"net/http"
	"strconv"
	"time"
)

// GetInfo 查询个人信息
//
//	@Summary		查询个人信息
//	@Description	会执行用户查询个人信息功能
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"输入token"
//	@Response		200,500			{object}	tools.HttpCode{data=model.APIUser}
//	@Router			/user/getInfo [get]
func GetInfo(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	var id int64
	var err error
	if id, err = tools2.GetIDFromToken(auth); err != nil {
		logger.Log.Error("从Token获取id失败")
		return
	}

	logger.Log.Error("已经获取id", id)
	if id <= 0 {
		logger.Log.Error("id数据有误")
		return
	}
	ret := &model2.APIUser{}
	if err := model2.GetInfo(ret, id); err != nil {
		c.JSON(http.StatusOK, tools2.HttpCode{
			Code:    tools2.UserInfoErr,
			Message: "获取用户信息成功失败",
			Data: struct {
			}{},
		})
		logger.Log.Error(err)
		return
	}

	c.JSON(http.StatusOK, tools2.HttpCode{
		Code:    tools2.OK,
		Message: "获取用户信息成功成功",
		Data:    ret,
	})
	return
}

// UpdateInfo 修改个人信息
//
//	@Summary		修改个人信息
//	@Description	会执行用户修改个人信息功能
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"输入token"
//	@Param			name			body		string	false	"用户名"
//	@Param			pwd				body		string	false	"密码"
//	@Param			tel				body		string	false	"手机号"
//	@Response		200,500			{object}	tools.HttpCode{data=model.APIUser}
//	@Router			/user/updateInfo [put]
func UpdateInfo(c *gin.Context) {
	ret := &model2.APIUser{}
	if err := c.ShouldBind(&ret); err != nil {
		c.JSON(http.StatusOK, tools2.HttpCode{
			Code:    tools2.UserInfoErr,
			Message: "数据绑定失败",
			Data: struct {
			}{},
		})
		logger.Log.Error("数据绑定失败")
		return
	}
	auth := c.GetHeader("Authorization")
	var id int64
	var err error
	if id, err = tools2.GetIDFromToken(auth); err != nil {
		logger.Log.Error("从Token获取id失败")
		return
	}

	logger.Log.Error("已经获取id", id)
	if id <= 0 {
		logger.Log.Error("id数据有误")
		return
	}

	if model2.CheckGetUserExist(ret.Name) == true {
		c.JSON(http.StatusOK, tools2.HttpCode{
			Code:    tools2.UserInfoErr,
			Message: "用户名已被占用,修改失败",
			Data: struct {
			}{},
		})
		logger.Log.Error("用户名已被占用,修改失败")
		return
	}

	if err := model2.UpdateInfo(ret, id); err != nil {
		c.JSON(http.StatusOK, tools2.HttpCode{
			Code:    tools2.UserInfoErr,
			Message: "修改用户信息成功失败",
			Data: struct {
			}{},
		})
		logger.Log.Error(err)
		return
	}

	c.JSON(http.StatusOK, tools2.HttpCode{
		Code:    tools2.OK,
		Message: "修改用户信息成功成功",
		Data:    ret,
	})
	return
}

// UserLendBook
//
//	@Summary		借书
//	@Description	用户借书籍
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"输入token"
//	@Param			id				path		int		false	"id"
//	@response		200,500			{object}	tools.HttpCode
//	@Router			/user/lendBook/{id} [put]
func UserLendBook(c *gin.Context) {
	bookIdStr := c.Param("id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	auth := c.GetHeader("Authorization")
	var id int64
	var err error
	if id, err = tools2.GetIDFromToken(auth); err != nil {
		logger.Log.Error("从Token获取id失败")
		return
	}

	logger.Log.Error("已经获取id", id)
	if id <= 0 {
		logger.Log.Error("id数据有误")
		return
	}

	ret := &model2.UserBook{}
	ret.LendTime = time.Now()

	if model2.UserLendBook(bookId, id) == false {
		c.JSON(http.StatusOK, tools2.HttpCode{
			Code:    tools2.UserInfoErr,
			Message: "借书失败",
			Data: struct {
			}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools2.HttpCode{
		Code:    tools2.OK,
		Message: "借书成功",
		Data: struct {
		}{},
	})
	return
}

// UserGiveBook
//
//	@Summary		还书
//	@Description	用户还书籍
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"输入token"
//	@Param			id				path		int		false	"id"
//	@response		200,500			{object}	tools.HttpCode
//	@Router			/user/giveBook/{id} [put]
func UserGiveBook(c *gin.Context) {
	bookIdStr := c.Param("id")
	bookId, _ := strconv.ParseInt(bookIdStr, 10, 64)
	auth := c.GetHeader("Authorization")
	var id int64
	var err error
	if id, err = tools2.GetIDFromToken(auth); err != nil {
		logger.Log.Error("从Token获取id失败")
		return
	}

	logger.Log.Info("已经获取用户id", id)
	if id <= 0 {
		logger.Log.Error("id数据有误")
		return
	}

	ret := &model2.UserBook{}
	ret.LendTime = time.Now()

	if model2.UserGiveBook(bookId, id) == false {
		c.JSON(http.StatusOK, tools2.HttpCode{
			Code:    tools2.UserInfoErr,
			Message: "还书失败",
			Data: struct {
			}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools2.HttpCode{
		Code:    tools2.OK,
		Message: "还书成功",
		Data: struct {
		}{},
	})
	return
}

// GetAllLendInfo
//
//	@Summary		个人借阅记录
//	@Description	会执行获得个人的借阅记录
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	false	"输入token"
//	@response		200,500			{object}	tools.HttpCode{data=model.APILendBooks}
//	@Router			/user/getAllLendInfo [get]
func GetAllLendInfo(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	var id int64
	var err error
	if id, err = tools2.GetIDFromToken(auth); err != nil {
		logger.Log.Error("从Token获取id失败")
		return
	}
	ret := &[]model2.APIUserBook{}
	if err = model2.GetAllLendInfo(ret, id); err != nil {
		c.JSON(http.StatusOK, tools2.HttpCode{
			Code:    tools2.UserInfoErr,
			Message: "查询个人借阅记录失败",
			Data: struct {
			}{},
		})
		logger.Log.Error(err)
		return
	}
	c.JSON(http.StatusOK, tools2.HttpCode{
		Code:    tools2.UserInfoErr,
		Message: "查询个人借阅记录成功",
		Data:    ret,
	})
	return
}
