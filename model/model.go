package model

import "time"

type User struct {
	Id          int64     `gorm:"id"`
	Name        string    `form:"name" binding:"required"` //用户名
	Pwd         string    `form:"pwd" binding:"required"`  //密码
	Tel         string    `form:"tel" binding:"required"`  //手机号
	CreatedTime time.Time //注册时间
}

type Admin struct {
	Id          int64     `gorm:"id"`
	Name        string    `form:"name" binding:"required"` //用户名
	Pwd         string    `form:"pwd" binding:"required"`  //密码
	Tel         string    `form:"tel" binding:"required"`  //手机号
	CreatedTime time.Time //注册时间
}

type Book struct {
	Id      int64  `gorm:"id"`
	Name    string `json:"name" binding:"required"`   //图书名字
	Author  string `json:"author" binding:"required"` //图书作者
	Number  string `json:"number" binding:"required"` //图书编号
	LendOut bool   `json:"lend_out"`                  //图书是否借出
	UserId  int64  `json:"user_id"`                   //借出用户
	//Count   int64  `json:"count"`                   //被借次数
}

type LendBooks struct {
	Id       int64     `gorm:"id"`
	BookId   int64     //图书Id
	BookName string    //图书名
	UserId   int64     //用户Id
	UserName string    //用户名
	LendTime time.Time //借出时间
	GiveBook bool      //是否归还
	GiveTime time.Time //归还时间
}

type APIUser struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
	Tel  string `json:"tel" form:"tel" binding:"required"`
}

type APIBook struct {
	Name    string
	Author  string
	Number  int64
	LendOut bool
}

type APILendBooks struct {
	BookName string
	UserName string
	LendTime time.Time
	GiveBook bool
	GiveTime time.Time
}
