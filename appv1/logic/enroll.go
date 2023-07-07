package logic

import (
	"Library-management/appv1/modle"
	"Library-management/appv1/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Enroll godoc
//
// @Summary		注册
// @Description	会执行注册操作
// @Tags			all
// @Accept			mpfd
// @Produce		  json
// @Param name formData string true "The name of the user"
// @Param password formData string true "The password for the user account"
// @Param phone formData string true "The phone number of the user"
// @Router			/all/enroll [POST]
func Enroll(c *gin.Context) {
	data := modle.User{}
	maps := make(map[string]interface{})
	maps["Name"] = c.PostForm("name")
	maps["Password"] = modle.Md(c.PostForm("password"))
	maps["Phone"] = c.PostForm("phone")
	modle.GlobalConn.Where("name = ?", maps["Name"]).First(&data)
	if data.UserId != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code":    tools.UserInfoErr,
			"Message": "用户已经存在",
		})
		return
	} else {
		ok := data.AddUser(maps)
		c.JSON(200, gin.H{
			"Code":    tools.OK,
			"Message": ok,
		})
	}
}
