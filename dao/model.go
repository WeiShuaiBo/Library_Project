package dao

import "time"

type User struct {
	Id         int64     `json:"id" gorm:"autoIncrement"` //用户ID
	Name       string    //用户名
	Pwd        string    //用户密码
	Type       bool      //用户的类别,0：普通用户,1：图书管理员
	CreateTime time.Time //用户创建时间
}
type Book struct {
	Id          int64     `json:"id" gorm:"autoIncrement"` //图书ID
	BookName    string    `json:"book_name"`               //书名
	Author      string    `json:"author"`                  //作者
	Description string    `json:"description"`             //描述
	Count       int       `json:"count"`                   //图书数量
	CreateTime  time.Time //图书添加时间
}
type UserBook struct {
	Id         int64     `json:"id"`          //序列号
	UId        int64     `json:"user_id"`     //用户ID
	BookId     int64     `json:"book_id"`     //图书ID
	Type       bool      `json:"type"`        // 图书借阅状态，0：借阅，1：已归还
	BorrowTime time.Time `json:"borrow_time"` //借阅的时间
	ReturnTime time.Time `json:"return_time"` //归还时间
}
type Xxx struct {
	id   int
	name string
}

func (User) TableName() string {
	return "user"
}
func (Book) TableName() string {
	return "book"
}
func (UserBook) TableName() string {
	return "user_book"
}
