package model

import (
	"library/appV0/logger"
)

func GetBooks(ret *[]APIBookInfo, limit, offset string) error {
	sql := "SELECT * FROM book_info LIMIT ? OFFSET ? "
	if err := GlobalConn.Raw(sql, limit, offset).Scan(&ret).Error; err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil
}
func GetBookByKeyWord(ret *[]APIBookInfo, keyWord string, limit, offset string) error {
	keyWord = "%" + keyWord + "%"
	sql := "SELECT * FROM book_info WHERE book_name LIKE ? LIMIT ? OFFSET ?"
	if err := GlobalConn.Raw(sql, keyWord, limit, offset).Scan(&ret).Error; err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil
}
