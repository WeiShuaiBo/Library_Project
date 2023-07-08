package dao

import (
	"fmt"
)

func GetAllUser(page int64) []User {
	var user []User
	sql := "select * from user limit 10 offset ?"
	err := GlobalConn.Debug().Raw(sql, (page-1)*10).Find(&user).Error
	if err != nil {
		fmt.Println("全部用户信息查询失败")
	}
	return user
}
func GetUserBorrow(id int64, page int64) []UserBook {
	var userBook []UserBook
	sql := "select * from user_book where user_id = ? limit 10 offset ?"
	err := GlobalConn.Debug().Raw(sql, id, (page-1)*10).Find(&userBook).Error
	if err != nil {
		fmt.Println("未找到指定用户的借阅信息")
	}
	return userBook
}
func GetALlBookBorrow(page int64) []UserBook {
	var userBook []UserBook
	sql := "select * from user_book limit 10 offset ?"
	err := GlobalConn.Raw(sql, (page-1)*10).Find(&userBook).Error
	if err != nil {
		fmt.Println("未找到全部图书的借阅信息")
	}
	return userBook
}
func AddBook(book Book) bool {
	var oldBook Book
	sql := "select * from book where id = ? limit 1"
	GlobalConn.Raw(sql, book.Id).Find(&oldBook)
	if oldBook.Id == book.Id {
		sql := "update book set count = count + ? where id = ?"
		if err := GlobalConn.Debug().Exec(sql, book.Count, book.Id).Error; err != nil {
			fmt.Println("添加图书数量失败")
			return false
		}
		return true
	}
	if err := GlobalConn.Debug().Table("book").Create(&book).Error; err != nil {
		fmt.Println("添加图书失败")
		return false
	}
	return true
}
func DeleteBook(bookId int64) bool {
	var book Book
	sql := "select * from book where id = ? limit 1"
	GlobalConn.Debug().Raw(sql, bookId).Find(&book)
	if book.Id < 1 {
		return false
	}
	var userBook []UserBook
	sql2 := "select * from user_book where book_id = ?"
	GlobalConn.Raw(sql2, bookId).Find(&userBook)
	if len(userBook) != 0 {
		for i, _ := range userBook {
			if userBook[i].Type != 1 {
				fmt.Println("该图书正在被借阅，无法删除")
				return false
			}
		}
	}
	sql1 := "update book set del_flg = 1 where id = ?"
	err := GlobalConn.Debug().Exec(sql1, bookId).Error
	if err != nil {
		fmt.Println("未能删除指定图书")
		return false
	}
	return true
}
