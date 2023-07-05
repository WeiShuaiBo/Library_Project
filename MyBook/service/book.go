// @Author	zhangjiaozhu 2023/7/4 10:27:00
package service

import (
	"MyBook/common/snowflake"
	"MyBook/dao"
	"MyBook/models"
	"MyBook/models/DB/MysqlDB"
	"errors"
	"strconv"
)

func GetBook() ([]*models.Book, error) {
	book, err := dao.GetBook()
	if err != nil {
		return nil, err
	}
	return book, nil
}

func FindBookByName(s string) (models.Book, error) {
	book, err := dao.FindBookByName(s)
	return book, err
}
func CreateBook(req *models.Book) error {
	id, err := snowflake.GetID()
	if err != nil {
		return errors.New("服务器繁忙")
	}
	req.BookId = id
	if err := dao.CreateBook(req); err != nil {
		return err
	}
	return nil
}
func DeleteBookByName(book models.Book) error {
	return dao.DeleteBookByName(book)
}

// 借书
func BorrowBook(bookName string, id uint64) error {

	// 通过图书名获取图书信息
	book, err := dao.FindBookByName(bookName)
	if err != nil {
		return err
	}
	// 判断图书库存是否大于0，大于0才能借书
	atoi, _ := strconv.Atoi(book.Remaining)
	if atoi <= 0 {
		return errors.New("图书库存不足，无法借书")
	}
	err = dao.BorrowBook(id, &book)
	if err != nil {
		return err
	}
	return nil
}
func ReturnBook(bookName string, id uint64) error {
	var book models.Book
	MysqlDB.DB.Where("book_name = ?", bookName).First(&book)
	err := dao.ReturnBook(id, &book)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(book *models.Book) error {
	return dao.UpdateBook(book)
}
