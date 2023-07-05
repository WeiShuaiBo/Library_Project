// @Author	zhangjiaozhu 2023/7/3 20:29:00
package dao

import (
	"MyBook/models"
	"MyBook/models/DB/MysqlDB"
	"errors"
	"gorm.io/gorm"
)

//func UserRegisterDao(user *models.UserRegister) (err error) {
//	sqlStr := "select count(user_id) from user where username = ?"
//	var count int64
//	err = MysqlDB.DB.Raw(sqlStr, user.UserName).Count(&count).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return err
//	}
//	if count > 0 {
//		// 用户已存在
//		return errors.New("用户已存在")
//	}
//	// 生成user_id
//	userID, err := snowflake.GetID()
//	if err != nil {
//		return errors.New("数据库繁忙")
//	}
//	// 生成加密密码
//	password := encryptPassword([]byte(user.Password))
//	// 把用户插入数据库
//	sqlStr = "insert into user(user_id, username, password) values (?,?,?)"
//	_, err = db.Exec(sqlStr, userID, user.UserName, password)
//	return
//}

func FindUserById(id uint64) (models.User, error) {
	var user models.User
	err := MysqlDB.DB.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return user, errors.New("服务器繁忙")
	}
	if user.UserId < 0 {
		return user, errors.New("查无此用户")
	}
	return user, nil
}

func UpdateUserInfo(user models.User) error {
	err := MysqlDB.DB.Model(models.User{}).Updates(user).Error
	if err != nil {
		return errors.New("系统繁忙，更新失败")
	}
	return nil
}

func DeleteUser(id uint64, user models.User) error {
	err := MysqlDB.DB.Where("user_id=?", id).Delete(&user).Error
	if err != nil {
		return errors.New("服务器繁忙")
	}
	return nil
}
func GetAllUser() ([]*models.User, error) {
	users := make([]*models.User, 10)
	err := MysqlDB.DB.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("服务器繁忙")
	}
	if len(users) == 0 {
		return nil, errors.New("空记录")
	}
	return users, nil
}
