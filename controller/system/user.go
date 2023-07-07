package system

import (
	"Library_Project/global"
	"Library_Project/model/system"
	"Library_Project/model/system/request"
	"Library_Project/pkg"
	"Library_Project/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SysUserController struct {
}

// Login 用户登录
// Login
// @Tags public
// @Summary 用户登录
// @Accept json
// @Produce json
// @Param data body request.Login true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/public/login [post]
func (sc *SysUserController) Login(c *gin.Context) {
	login := request.Login{}
	err := c.ShouldBindJSON(&login)
	if err != nil {
		global.FAST_LOG.Error("注册数据绑定错误:" + err.Error())
		c.JSON(200, "注册数据绑定错误"+err.Error())
		return
	}
	if service.ServiceApp.SystemServiceGroup.Login(login.UserId, login.Password) {

		token, err1 := pkg.GenToken(login.UserId)
		fmt.Println(token)
		fmt.Println(err1)
		if err1 != nil {
			c.JSON(200, "账号或密码错误")
			return
		}
		c.JSON(200, gin.H{
			"token": token,
		})

	}

}

// Register 用户注册
// @Tags public
// @Summary 用户注册账号
// @Accept json
// @Produce json
// @Param data body request.Login true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/public/register [post]
func (sc *SysUserController) Register(c *gin.Context) {
	login := request.Login{}
	err := c.ShouldBindJSON(&login)
	if err != nil {
		global.FAST_LOG.Error("注册数据绑定错误:" + err.Error())
		c.JSON(200, "注册数据绑定错误"+err.Error())
		return
	}
	if !service.ServiceApp.SystemServiceGroup.ExistId(login.UserId) {
		service.ServiceApp.SystemServiceGroup.InsertUser(login.UserId, login.Password)
		c.JSON(200, "注册成功")
		return
	}
	c.JSON(200, "账号存在")

}

// Information 用户信息
// @Tags user
// @Summary 获取用户信息
// @Accept json
// @Produce json
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/information [get]
func (sc *SysUserController) Information(c *gin.Context) {

	value := c.GetInt("UserId")
	user := service.ServiceApp.SystemServiceGroup.Information(value)

	c.JSON(200, gin.H{
		"code": 200,
		"date": user,
		"msg":  "个人信息",
	})

}

// ChangeInfo 修改用户信息
// @Tags user
// @Summary 修改用户信息
// @Accept json
// @Produce json
// @Param data body system.User true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/changeinfo [post]
func (sc *SysUserController) ChangeInfo(c *gin.Context) {
	userid := c.GetInt("UserId")
	u := &system.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"date": "",
			"mgs":  "修改失败",
		})
		return
	}
	service.ServiceApp.SystemServiceGroup.ChangeInfo(userid, u)
	c.JSON(200, gin.H{
		"code": 200,
		"date": u,
		"mgs":  "修改成功",
	})
}

// Borrow 借书
// @Tags user
// @Summary 借书
// @Accept json
// @Produce json
// @Param data body system.Book true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/borrow [post]
func (sc *SysUserController) Borrow(c *gin.Context) {
	book := &system.Book{}
	err := c.ShouldBindJSON(&book)
	if err != nil {
		global.FAST_LOG.Error("借书出错：" + err.Error())
		c.JSON(200, gin.H{
			"code": 200,
			"date": "",
			"mgs":  "借书失败",
		})
		return
	}
	fmt.Println(book)

	userid := c.GetInt("UserId")
	borrow := service.ServiceApp.SystemServiceGroup.Borrow(userid, book)
	if borrow == 1 {
		c.JSON(200, gin.H{
			"code": 200,
			"date": "",
			"msg":  "借书成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"date": "",
			"msg":  "不能借阅",
		})
	}

}

// ReturnBook 还书
// @Tags user
// @Summary 还书
// @Accept json
// @Produce json
// @Param data body system.Book true "body参数"
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/returnbook [post]
func (sc *SysUserController) ReturnBook(c *gin.Context) {
	book := &system.Book{}
	err := c.ShouldBind(&book)
	if err != nil {
		global.FAST_LOG.Error("还书出错：" + err.Error())
		c.JSON(200, gin.H{
			"code": 200,
			"date": "",
			"mgs":  "还书失败",
		})
		return
	}
	re := service.ServiceApp.SystemServiceGroup.Return(book)
	if re == 0 {
		c.JSON(200, gin.H{
			"code": 200,
			"date": "",
			"msg":  "还书失败，逾期了",
		})
	} else if re == 1 {
		c.JSON(200, gin.H{
			"code": 200,
			"date": "",
			"msg":  "还书成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"date": "",
			"msg":  "书籍未借出，不能还书",
		})
	}
}

// Exit 退出登录
// @Tags user
// @Summary 退出登录
// @Accept json
// @Produce json
// @Success 200 {string} json "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /api/sys/exit [get]
func (sc *SysUserController) Exit(c *gin.Context) {
	c.Request.Header.Set("Authorization", "")
	c.JSON(200, gin.H{
		"code": 200,
		"date": "",
		"msg":  "退出",
	})
}
