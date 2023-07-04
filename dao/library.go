package dao

import "fmt"

func GetBook(bookName string) Book {
	var book Book
	sql := "select * from book where book_name = ?"
	if err := GlobalConn.Debug().Raw(sql, bookName).Find(&book).Error; err != nil {
		fmt.Println("指定图书查询错误")
	}
	return book
}
func GetAllBook() []Book {
	var book []Book
	sql := "select * from book"
	if err := GlobalConn.Raw(sql).Find(&book).Error; err != nil {
		fmt.Println("全部图书内容查询失败")
	}
	return book
}
