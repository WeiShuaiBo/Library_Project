package model

import (
	"errors"
	"library/appV0/logger"
)

func AdminGetBooks(ret *[]BookInfo, limit, offset string) error {
	sql := "SELECT * FROM book_info LIMIT ? OFFSET ? "
	if err := GlobalConn.Raw(sql, limit, offset).Scan(&ret).Error; err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil
}

func AdminGetBookByKeyWord(ret *[]BookInfo, keyWord string, limit, offset string) error {
	keyWord = "%" + keyWord + "%"
	sql := "SELECT * FROM book_info WHERE book_name LIKE ? LIMIT ? OFFSET ?"
	if err := GlobalConn.Raw(sql, keyWord, limit, offset).Scan(&ret).Error; err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil
}

func AdminGetBooksById(ret *BookInfo, id, limit, offset string) error {
	sql := "SELECT * FROM book_info WHERE id = ? LIMIT ? OFFSET ?"
	if err := GlobalConn.Raw(sql, id, limit, offset).Scan(&ret).Error; err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil
}

func CreatBook(book *BookInfo) error {
	sql := "insert into book_info (book_name,author,isbn) values (?,?,?)"
	if err := GlobalConn.Exec(sql, book.BookName, book.Author, book.ISBN).Error; err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil
}

func UpdateBook(book *BookInfo, isbnStr string) error {
	sql := "update book_info set book_name=?,author=?,isbn=? where isbn=?"
	if err := GlobalConn.Exec(sql, book.BookName, book.Author, book.ISBN, isbnStr).Error; err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil
}

func DeleteBook(isbnStr string) error {
	sql := "delete from book_info where isbn=?"
	if GlobalConn.Exec(sql, isbnStr).RowsAffected <= 0 {
		logger.Log.Error("删除失败")
		return errors.New("删除失败")
	}
	return nil
}

func AdminGetInfo(ret *[]UserBook, limit, offset string) error {
	sql := "SELECT * FROM user_book LIMIT ? OFFSET ?"
	if err := GlobalConn.Raw(sql, limit, offset).Scan(&ret).Error; err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil
}
