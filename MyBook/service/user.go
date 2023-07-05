// @Author	zhangjiaozhu 2023/7/3 19:33:00
package service

import (
	"MyBook/common/md5"
	"MyBook/common/snowflake"
	"MyBook/dao"
	"MyBook/models"
	"MyBook/models/DB/MysqlDB"
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func UserRegister(user *models.UserRegister) error {
	var count int64
	if err := MysqlDB.DB.Raw(`select count(user_id) from user where user_name = ?`, user.UserName).Scan(&count).Debug().Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	// 生成UUID
	userId, err := snowflake.GetID()
	if err != nil {
		return err
	}
	// 加密密码
	password := md5.EncryptPassword([]byte(user.Password))
	// 插入数据库
	var userModel models.User
	userModel.UserName = user.UserName
	userModel.Password = password
	userModel.UserId = userId
	userModel.Email = user.Email
	userModel.Role = 0

	err = MysqlDB.DB.Create(&userModel).Debug().Error
	//err = MysqlDB.DB.Exec(`insert into user(user_id, username, password,email,role) values (?,?,?,?,?)`, userId, user.UserName, password, user.Email, "0").Debug().Error
	if err != nil {
		return err
	}
	return nil
}

func UserLogin(u models.UserLogin) (models.User, error) {
	orignPassword := u.Password // 记录原始密码
	name := u.UserName

	var user models.User
	sqlStr := `select user_id, user_name, password from user where user_name = ?`
	err := MysqlDB.DB.Raw(sqlStr, name).Scan(&user).Debug().Error

	if err != nil && err != gorm.ErrRecordNotFound {
		// 查询数据库出错
		return models.User{}, err
	}
	if err == gorm.ErrRecordNotFound {
		// 用户不存在
		return models.User{}, errors.New("用户不存在")
	}
	// 生成加密密码与查询到的密码比较
	password := md5.EncryptPassword([]byte(orignPassword))
	if user.Password != password {
		return models.User{}, errors.New("用户密码不正确")
	}
	return user, nil
}
func FindUserById(id uint64) (models.User, error) {
	return dao.FindUserById(id)
}

func UpdateUserInfo(user models.User) error {
	return dao.UpdateUserInfo(user)
}

func DeleteUser(id string) error {
	atoi, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("类型转换失败")
	}
	user, err := dao.FindUserById(uint64(atoi))
	if err != nil {
		return err
	}
	if user.UserId == 0 {
		return errors.New("用户不存在")
	}
	return dao.DeleteUser(uint64(atoi), user)
}

func GetAllUser() ([]*models.User, error) {
	return dao.GetAllUser()
}
