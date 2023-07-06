package dao

import (
	"fmt"
	"time"
)

func CheckUser(name, pwd string) User {
	var user User
	sql := "select * from user where name = ? and pwd = ?"
	err := GlobalConn.Raw(sql, name, pwd).Find(&user).Error
	if err != nil {
		fmt.Println("未找到登录用户信息")
	}
	return user
}
func Register(user User) error {
	//sql := "insert into user(name,pwd,type,create_time) values(?,?,?,?)"
	//err := GlobalConn.Debug().Raw(sql, user.Name, user.Pwd, user.Type, user.CreateTime).Error
	err := GlobalConn.Debug().Table("user").Create(&user).Error
	return err
}
func GetUser(id int) User {
	sql := "select * from user where id = ? limit 1"
	var user User
	if err := GlobalConn.Debug().Raw(sql, id).Find(&user).Error; err != nil {
		fmt.Println("用户信息查询失败")
	}
	return user
}
func PutUser(user *User) error {
	return GlobalConn.Where("id = ?", user.Id).Updates(user).Error
}
func GetMyBook(id int64) []UserBook {
	var recode []UserBook
	sql := "select * form user_book where id = ?"
	if err := GlobalConn.Raw(sql, id).Find(&recode).Error; err != nil {
		fmt.Println("查找个人借阅信息失败")
	}
	return recode
}
func Borrow(id int64, bookName string) bool {
	var book Book
	sql := "select * from book where book_name = ? limit 1"
	if err := GlobalConn.Raw(sql, bookName).Find(&book).Error; err != nil {
		fmt.Println("未找到指定图书")
		return false
	}
	if book.Count < 1 {
		fmt.Println("图书数量不足")
		return false
	}
	tx := GlobalConn.Begin()
	sql1 := "update book set count = count -1 where name = ?"
	if err := tx.Table("book").Raw(sql1, bookName).Error; err != nil {
		fmt.Println("借阅图书操作失败")
		tx.Rollback()
		return false
	}
	var userBook = UserBook{
		UserId:     id,
		BookId:     book.Id,
		Type:       0,
		BorrowTime: time.Now(),
	}
	if err := tx.Table("user_book").Create(&userBook).Error; err != nil {
		fmt.Println("借阅图书操作失败")
		tx.Rollback()
		return false
	}
	return true
}
func GiveBack(id int64, bookName string) bool {
	var book Book
	sql := "select * from book where book_name = ? limit 1"
	if err := GlobalConn.Raw(sql, bookName).Find(&book).Error; err != nil {
		fmt.Println("未找到指定图书")
		return false
	}
	tx := GlobalConn.Begin()
	sql1 := "update book set count = count + 1 where name = ?"
	if err := tx.Table("book").Raw(sql1, bookName).Error; err != nil {
		fmt.Println("归还图书操作失败")
		tx.Rollback()
		return false
	}
	m := map[string]interface{}{
		"type":        1,
		"return_time": time.Now(),
	}
	if err := tx.Table("user_book").Where("book_id = ? and user_id = ?", book.Id, id).Updates(m).Error; err != nil {
		fmt.Println("归还图书操作失败")
		tx.Rollback()
		return false
	}
	return true
}
