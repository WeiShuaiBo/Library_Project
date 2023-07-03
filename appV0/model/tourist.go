package model

import (
	"fmt"
	"strconv"
)

func GetAll() []*Book {
	book := make([]*Book, 0)
	sql := "select * from book LIMIT 100"
	err := GlobalConn.Raw(sql).Scan(&book).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return book
	}

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
