package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/appV0/model"
	"library/appV0/tools"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// BorrowBook godoc
// 接口的名字
//
//		@Summary		用户借书
//		@Description	会执行根据bookId借书操作
//		@Tags			user
//		@Accept			multipart/form-data
//		@Produce		json
//	 @Param			book[]	formData	[]int	false	"图书ID"
//
// @Param Authorization header string false "Bearer 用户令牌"
//
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/user/users/records/{bookId}
func BorrowBook(c *gin.Context) {
	userId := int64(0) //user_id 为0  说明是测试用例
	if uid, ok := c.Get("userId"); ok {
		userId = uid.(int64)
	}

	fmt.Print("userId:")
	fmt.Println(userId)

	bookId, _ := c.GetPostFormArray("book[]")
	bookIds := make([]int64, 0)
	slice := strings.Join(bookId, ",")
	bookId = strings.Split(slice, ",")
	fmt.Println("bookid++++++++++++++++++")
	fmt.Println("Type:", reflect.TypeOf(bookId))
	for _, val := range bookId {
		tmp, _ := strconv.ParseInt(val, 10, 64)
		bookIds = append(bookIds, tmp)
		fmt.Printf("____bookId___")
		fmt.Println(tmp)
	}
	fmt.Println("bookids++++++++++++++++++")
	fmt.Print(bookIds)
	//record := &model.Record{}

	if model.BorrowBook(userId, bookIds) {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "借书成功",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.UserInfoErr,
		Message: "所借图书已被借走",
		Data:    struct{}{},
	})
	return
}

// ReturnBook godoc
// 接口的名字
//
//	@Summary		还书
//	@Description	会执行根据bookId还书操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	 @Param			book[]	formData	[]int	false	"图书ID"
//
// @Param Authorization header string false "Bearer 用户令牌"
//
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/user/users/records/{bookId} [Put]
func ReturnBook(c *gin.Context) {
	userId := int64(0) //user_id 为0  说明是测试用例
	if uid, ok := c.Get("userId"); ok {
		userId = uid.(int64)
	}

	fmt.Print("userId:")
	fmt.Println(userId)

	bookId, _ := c.GetPostFormArray("book[]")
	bookIds := make([]int64, 0)
	slice := strings.Join(bookId, ",")
	bookId = strings.Split(slice, ",")
	fmt.Println("bookid+++++++++++++")
	fmt.Println("Type:", reflect.TypeOf(bookId))
	for _, val := range bookId {
		tmp, _ := strconv.ParseInt(val, 10, 64)
		bookIds = append(bookIds, tmp)
		fmt.Printf("____bookId___")
		fmt.Println(tmp)
	}
	fmt.Println("bookids+++++++++++")
	fmt.Println(bookIds)
	//record := &model.Record{}

	if model.ReturnBook(userId, bookIds) {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.OK,
			Message: "还书成功",
			Data:    struct{}{},
		})
		return
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.UserInfoErr,
		Message: "代换书籍有误",
		Data:    struct{}{},
	})
	return
}

// UserLogin godoc
// 接口的名字
//
//	@Summary		用户登录
//	@Description	会执行用户登录操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name	formData	string	true	"用户名"
//	@Param			pwd		formData	string	true	"密码"
//
//
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/userLogin [POST]
func UserLogin(c *gin.Context) {
	//声明user
	var user User
	//	绑定并判断
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "绑定失败",
			Data:    struct{}{},
		})
		return
	}
	fmt.Printf("看看user")
	fmt.Print(user)
	//	查询数据库并判断
	dbUser := model.GetUser(user.Name, user.Pwd)
	if dbUser.Id <= 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "查询数据库失败",
			Data:    struct{}{},
		})
		return
	}
	//	下发Token并判断
	a, r, err := tools.Token.GetToken(dbUser.Id, dbUser.Name)
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

// Logout godoc
//
//	@Summary		用户退出
//	@Description	会执行用户退出操作
//	@Tags			user
//	@Accept			multipart/form-data
//
// @Param Authorization header string false "Bearer 用户令牌"
//
//	@Produce		json
//	@response		500,401	{object}	tools.HttpCode
//	@Router			/userLogout [get]
func Logout(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, tools.HttpCode{
		Code: tools.OK,
		Data: struct{}{},
	})
	return
}
