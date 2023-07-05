// @Author	zhangjiaozhu 2023/7/3 19:15:00
package models

import (
	"MyBook/models/DB/MysqlDB"
	"gorm.io/gorm"
)

type UserLogin struct {
	UserName string `form:"username" json:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Code     string `json:"code" form:"code"`
}

type UserRegister struct {
	UserName string `form:"username" json:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Code     string `json:"code" form:"code"`
}

type User struct {
	gorm.Model
	UserName string
	Password string
	Email    string
	UserId   uint64
	Role     int
}

type UserUpdate struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *User) TableName() string {
	return "user"
}

func (*UserRegister) Creat(user *UserRegister) error {
	if err := MysqlDB.DB.Create(user).Debug().Error; err != nil {
		return err
	}
	return nil
}
func (*UserRegister) SelectUserList(userList []*UserRegister) error {
	if err := MysqlDB.DB.Find(userList).Debug().Error; err != nil {
		return err
	}
	return nil
}
func (*UserRegister) GetUserById(id int64) (user *UserRegister, err error) {
	if err := MysqlDB.DB.Where("id=?", id).First(user).Debug().Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (*UserRegister) UserIsExist(name string) bool {
	if err := MysqlDB.DB.Where("username=?", name).Debug().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return true
		}
	}
	return true
}
func (*UserRegister) UpdateUser(user *UserRegister) (err error) {
	if err := MysqlDB.DB.Save(user).Debug().Error; err != nil {
		return err
	}
	return nil
}
