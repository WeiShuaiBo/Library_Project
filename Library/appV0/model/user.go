package model

import (
	"fmt"
	"time"
)

// GetUser 根据用户名和密码登录
func GetUser(name, pwd string) *User {

	ret := &User{}
	//sql语句：SELECT * FROM users WHERE name = ? and pwd = ?
	if err := GlobalConn.Table("users").Where("name = ? and pwd = ?", name, pwd).First(ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}

// FindUserName 查找用户名
func FindUserName(name string) bool {

	ret := &User{}
	//sql语言：SELECT * FROM users WHERE name = ?
	if err := GlobalConn.Table("users").Where("name=?", name).First(ret).Error; err != nil {
		return false
	}
	return true
}

// Register 注册
func Register(name, pwd, tel string) bool {

	ret := &User{
		Name:        name,
		Pwd:         pwd,
		Tel:         tel,
		CreatedTime: time.Now(),
		Privilege:   0,
	}

	//sql语言：INSERT INTO user(name.pwd,tel,createdTime) values(?,?,?,?)
	if err := GlobalConn.Table("users").Create(ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return false
	}
	return true
}

// UpdateUserPwd 修改用户信息
func UpdateUserPwd(user *User) error {

	//sql语言：UPDATE user set pwd =? where name = ?
	return GlobalConn.Table("users").Where("name=?", user.Name).Updates(user).Error
}

// FindUserPwdByName 查找用户名
func FindUserPwdByName(name string) *User {
	ret := &User{}

	//sql语言：SELECT * FROM users WHERE name = ?
	if err := GlobalConn.Table("users").Where("name=?", name).First(ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}

// FindUserPrivilegeById 查询用户权限
func FindUserPrivilegeById(id int64) int {
	ret := &User{}
	//sql := "SELECT privilege FROM users WHERE id = ?"
	//if err := GlobalConn.Table("users").Exec(sql, id).Find(ret).Error; err != nil {
	//	fmt.Printf("err:%s\n", err.Error())
	//}

	if err := GlobalConn.Table("users").Where("id = ?", id).Find(ret).Error; err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}

	fmt.Println("ret", ret)
	return ret.Privilege
}
