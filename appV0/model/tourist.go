package model

import (
	"fmt"
	"strconv"
)

func GetAll(pageNumber int) []*Book {
	book := make([]*Book, 0)
	pageSize := 10 //每页显示的数据条数。
	//pageNumber := pageNumber                       //要查询的页码。
	offset := (pageNumber - 1) * pageSize //偏移量，用于确定从数据库中的哪一行开始获取数据。

	var total int
	result := GlobalConn.Raw("SELECT COUNT(*) FROM book").Scan(&total)
	if result.Error != nil {
		fmt.Println("Failed to count books:", result.Error)
		return book
	}

	if pageNumber > int(total/pageSize)+1 {
		fmt.Println("Invalid page number")
		return book
	}

	sql := "select * from book LIMIT ? OFFSET ?"
	err := GlobalConn.Raw(sql, pageSize, offset).Scan(&book).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		fmt.Println("err！=nil")
		return book
	}
	// 统计总记录数
	//var total int64
	//result = db.Raw("SELECT COUNT(*) FROM users").Scan(&total)
	//if result.Error != nil {
	//	fmt.Println("Failed to count users:", result.Error)
	//	return
	//}
	//
	//fmt.Println("Total records:", total)

	return book
}

func GetBook(id int64) *Book {
	var book *Book

	sql := "select * from book where id = " + strconv.FormatInt(id, 10)

	err := GlobalConn.Raw(sql).Scan(&book).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return book
	}
	return book
}
func GetBookByName(bookName string) *BookInfo {
	var book *BookInfo
	fmt.Printf("1111111112111")
	fmt.Print(bookName)
	sql := "select * from book_info where book_name LIKE ?"

	err := GlobalConn.Raw(sql, "%"+bookName+"%").Scan(&book).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return book
	}
	return book
}
