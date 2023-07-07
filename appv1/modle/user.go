package modle

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUser(name, password string) *User {
	ret := &User{}
	if err := GlobalConn.Table("users").Where("name = ? and password = ?", name, password).First(ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}
func Md(s string) string {
	b := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", b)
}
func (u *User) AddUser(data map[string]interface{}) bool {
	user := User{}

	if name, ok := data["Name"].(string); ok {
		user.Name = name
	}

	if password, ok := data["Password"].(string); ok {
		user.Password = password
	}

	if phone, ok := data["Phone"].(string); ok {
		user.Phone = phone
	}

	GlobalConn.Create(&user)
	return true
}

// GetCurrentUserID
func GetCurrentUserID(c *gin.Context) int64 {

	userID := 33

	return int64(userID)
}

func UpdataUser(user *User) error {
	oldUser := User{}
	err := GlobalConn.Table("users").Where("user_id=?", user.UserId).First(&oldUser).Error
	if err != nil || (oldUser.UserId != 0 && oldUser.UserId != user.UserId) {
		return errors.New("查询失败")
	}
	return GlobalConn.Where("name=? AND password=? AND phone=?", user.Name, user.Password, user.Phone).Updates(user).Error
}
