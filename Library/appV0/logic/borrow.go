package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/appV0/model"
	"library/appV0/tools"
	"net/http"
	"strconv"
	"time"
)

// LoanLibrary godoc
//
//	@Summary		借书图书
//	@Description	根据ID借书图书
//	@Tags			user/common
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			library_id		query	string	true	"图书Id"
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,404,500	{object}	tools.HttpCode{data=Token}
//	@Router			/user/common/loanLibrary [POST]
func LoanLibrary(c *gin.Context) {
	userId := int64(0)
	ui, exists := c.Get("userId")
	if exists {
		userId = ui.(int64)
	}

	fmt.Println("userId:", userId)

	borrow1 := model.FindRecordByUserId(userId)

	fmt.Println("borrow1:", borrow1)

	if borrow1.UserId > 0 && !borrow1.IsReturn {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.DoErr,
			Message: "你借的书还未回还,截止时间为：",
			Data:    borrow1.DueDate,
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

	// 获取当前日期并格式化为字符串
	loanStr := time.Now().Format("2006-01-02")

	// 解析日期字符串为日期类型
	loanDate, _ := time.Parse("2006-01-02", loanStr)

	// 在日期上加上30天
	dueDate := loanDate.AddDate(0, 0, 30)

	// 打印结果
	fmt.Println("loanDate:", loanDate)
	fmt.Println("dueDate:", dueDate)

	borrow := &model.Borrow{
		UserId:    userId,
		LibraryId: libraryId,
		LoanDate:  time.Now(),
		DueDate:   dueDate,
		IsReturn:  false,
	}

	fmt.Println("borrow:", borrow)

	err1 := model.LoanLibrary(borrow)
	if err1 != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "借书失败",
			Data:    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "借书成功",
		Data:    struct{}{},
	})
}

// DueLibrary godoc
//
//	@Summary		归还图书
//	@Description	根据ID归还图书
//	@Tags			user/common
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id		query	string	true	"还书Id"
//	@Param Authorization header string false "Bearer 用户令牌"
//	@response		200,404,500	{object}	tools.HttpCode{data=Token}
//	@Router			/user/common/dueLibrary [POST]
func DueLibrary(c *gin.Context) {

	id := c.Query("id")
	// 将字符串转换为int64
	borrowId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Println("无法转换id为int64:", err)
		return
	}

	// 打印结果
	fmt.Println("libraryId:", borrowId)

	err = model.DueLibrary(borrowId)

	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.LibraryInfoErr,
			Message: "还书失败",
			Data:    err,
		})
		return
	}

	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "还书成功",
		Data:    struct{}{},
	})
}
