package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"library/appV0/SendEmail"
	"library/appV0/model"
	"library/appV0/tools"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CaptchaRequest struct {
	UserEmail string `json:"UserEmail"`
}
type User1 struct {
	UserName  string `form:"UserName" json:"UserName" binding:"required"`
	Password  string `form:"Password" json:"Password" binding:"required"`
	Name      string `form:"Name" json:"Name" binding:"required"`
	Sex       string `form:"Sex" json:"Sex" binding:"required"`
	Phone     string `form:"Phone" json:"Phone" binding:"required"`
	UserEmail string `form:"UserEmail" json:"UserEmail" binding:"required"`
}

// BorrowBook godoc
//
//	@Summary		借书
//	@Description	会执行根据bookId借书操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	 @Param			book[]	formData	[]int	false	"图书ID"
//
// @Param Authorization header string false "Bearer 用户令牌"
//
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/user/users/records/{bookId} [POST]
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
//	@Router			/user/users/records/{bookId} [PUT]
func ReturnBook(c *gin.Context) {
	fmt.Println("00000000000000000")
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
	fmt.Println(dbUser.Id)
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

// Captcha  godoc
//
//	@Summary		请求验证码
//	@Description	请求发送验证码
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			UserEmail	formData	string	true	"电子邮箱"
//	@response		500,401	{object}	tools.HttpCode
//	@Router			/Captcha [POST]
func Captcha(c *gin.Context) {
	var request CaptchaRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(c.Request.Body)
	toUserEmail := strings.Replace(request.UserEmail, "-", "", -1)
	code := generateCode()
	fmt.Print("向")
	fmt.Printf(toUserEmail)
	fmt.Print("发送：")
	fmt.Println(code)
	err := SendEmail.SendEmail(toUserEmail, code)

	if err != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "你配拥有验证码吗？",
			Data:    struct{}{},
		})
		return
	}
	client := model.CreateRedisClient()
	errRedis := client.Set(toUserEmail, code, 0).Err()
	if errRedis != nil {
		// 处理存储出错的情况
	}
	c.JSON(http.StatusOK, tools.HttpCode{
		Code:    tools.OK,
		Message: "验证码已发送",
		Data:    struct{}{},
	})
	return
}
func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999)
	return fmt.Sprintf("%06d", code)
}

// Registered godoc
//
//	@Summary		注册
//	@Description	会执行用户注册操作
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			UserName	formData	string	true	"用户名"
//	@Param			Password		formData	string	true	"密码"
//	@Param			Name		formData	string	true	"名字"
//	@Param			Sex		formData	string	true	"性别"
//	@Param			Phone		formData	string	true	"手机号"
//	@Param			UserEmail	formData	string	true	"电子邮箱"
//	@Param			code			path		int	true	"int valid"	minimum(1)
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//	@Router			/Registered/{code} [POST]
func Registered(c *gin.Context) {
	codeNew := c.Param("code")
	user := &User1{}
	//	绑定并判断
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "绑定失败",
			Data:    struct{}{},
		})
		return
	}
	client := model.CreateRedisClient()
	code, err := client.Get(user.UserEmail).Result()
	if err == redis.Nil {
		// 键不存在的处理逻辑
	} else if err != nil {
		// 获取代码出错的处理逻辑
	} else {
		// 使用获取到的代码进行后续操作
	}
	fmt.Println(code)
	fmt.Println(codeNew)
	if code != codeNew {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "验证码错误",
			Data:    struct{}{},
		})
		return
	}

	ret := model.Registered(user.UserName, user.Password, user.Name, user.Sex, user.Phone)
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

// GetPersonalInformation godoc
// @Summary 查看个人信息
// @Description 根据用户ID获取用户信息
// @Tags user
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {object} tools.HttpCode
// @Router /user/users/GetPersonalInformation/ [get]
func GetPersonalInformation(c *gin.Context) {
	userId := int64(0)
	if uid, ok := c.Get("userId"); ok {
		userId = uid.(int64)
	}
	ret := model.GetUserInformation(userId)
	if ret.Id > 0 {
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

// UpdatePersonalInformation godoc
//
//	@Summary		修改个人信息
//	@Description	会执行修改个人信息
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			UserName	formData	string	true	"用户名"
//	@Param			Password		formData	string	true	"密码"
//	@Param			Name		formData	string	true	"名字"
//	@Param			Sex		formData	string	true	"性别"
//	@Param			Phone		formData	string	true	"手机号"
//	@Param			UserEmail	formData	string	true	"电子邮箱"
//
// @Param Authorization header string false "Bearer 用户令牌"
//
//	@response		200,500	{object}	tools.HttpCode{data=Token}
//
// @Router 			/user/users/UpdatePersonalInformation/ [POSt]
func UpdatePersonalInformation(c *gin.Context) {
	user := &User1{}
	//	绑定并判断
	if c.ShouldBind(&user) != nil {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "绑定失败",
			Data:    struct{}{},
		})
		return
	}
	n := model.UpdatePersonalInformation(user.UserName, user.Password, user.Name, user.Sex, user.Phone)
	if n == 1 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "数据库出现问题",
			Data:    struct{}{},
		})
		return
	}
	if n == 2 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "没有找到对应username",
			Data:    struct{}{},
		})
		return
	}
	if n == 0 {
		c.JSON(http.StatusOK, tools.HttpCode{
			Code:    tools.UserInfoErr,
			Message: "修改成功",
			Data:    struct{}{},
		})
		return
	}
}
