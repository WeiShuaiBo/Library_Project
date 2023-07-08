package dao

import "fmt"

func GetBook(bookName string, page int64) []Book {
	var book []Book
	bookName = "%" + bookName + "%"
	sql1 := "select * from book where book_name like ? limit 10 offset ?"
	err := GlobalConn.Debug().Raw(sql1, bookName, (page-1)*10).Find(&book).Error
	if err != nil {
		fmt.Println("未能找到相关图书")
	}
	return book
}
func GetAllBook(page int64) []Book {
	var book []Book
	sql := "select * from book limit 10 offset ?"
	if err := GlobalConn.Raw(sql, (page-1)*10).Find(&book).Error; err != nil {
		fmt.Println("全部图书内容查询失败")
	}
	return book
}
