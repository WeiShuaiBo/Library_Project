package dao

import (
	"fmt"
)

func GetAllUser() []User {
	var user []User
	sql := "select * from user"
	err := GlobalConn.Raw(sql).Find(&user).Error
	if err != nil {
		fmt.Println("全部用户信息查询失败")
	}
	return user
}
func GetUserBorrow(id int64) []UserBook {
	var userBook []UserBook
	sql := "select * from user_book where user_id = ?"
	err := GlobalConn.Raw(sql, id).Find(&userBook).Error
	if err != nil {
		fmt.Println("未找到指定用户的借阅信息")
	}
	return userBook
}
func GetALlBookBorrow() []UserBook {
	var userBook []UserBook
	sql := "select * from user_book"
	err := GlobalConn.Raw(sql).Find(&userBook).Error
	if err != nil {
		fmt.Println("未找到全部图书的借阅信息")
	}
	return userBook
}
func AddBook(book Book) error {
	//sql := "insert into user(name,pwd,type,create_time) values(?,?,?,?)"
	//err := GlobalConn.Debug().Raw(sql, user.Name, user.Pwd, user.Type, user.CreateTime).Error
	err := GlobalConn.Debug().Table("book").Create(&book).Error
	return err
}
func DeleteBook(bookName string) bool {
	//sql := "delete from book where book_name = ?"
	//err := GlobalConn.Debug().Raw(sql, bookName).Error
	err := GlobalConn.Debug().Table("book").Where("book_name = ?", bookName).Delete(&Book{}).Error
	fmt.Println(err)
	if err != nil {
		fmt.Println("未能删除指定图书")
		return false
	}
	return true
}
