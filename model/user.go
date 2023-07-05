package model

import (
	"library/logger"
	"time"
)

func GetUser(name, pwd string) *User {
	ret := &User{}
	if err := GlobalConn.Table("users").Where("name = ? and pwd = ?", name, pwd).First(ret).Error; err != nil {
		logger.Log.Error("查询错误", err.Error())
	}
	return ret
}

func GetAdmin(name, pwd string) *User {
	ret := &User{}
	if err := GlobalConn.Table("admins").Where("name = ? and pwd = ?", name, pwd).First(ret).Error; err != nil {
		logger.Log.Error("查询错误", err.Error())
	}
	return ret
}

func CheckGetUserExist(name string) bool {
	ret := &User{}
	if GlobalConn.Table("users").Where("name = ?", name).First(ret).RowsAffected <= 0 {
		return false
	}

	return true
}

func RegisterUser(u *User) error {
	return GlobalConn.Create(u).Error
}

func GetInfo(ret *APIUser, id int64) error {
	return GlobalConn.Table("users").Where("id=?", id).First(&ret).Error
}

func UpdateInfo(ret *APIUser, id int64) error {
	return GlobalConn.Table("users").Where("id=?", id).Update(&ret).Error
}

//借书

func UserLendBook(bookId, UserId int64) bool {
	var err error
	tx := GlobalConn.Begin()
	if bookIsLend(bookId, true) == true {
		logger.Log.Error("该书已经被借出err:")
		return false
	}
	err = tx.Table("books").Where("id=?", bookId).
		Update("lend_out", true).Error
	if err != nil {
		logger.Log.Error(err)
		tx.Rollback()
		return false
	}

	ret := &LendBooks{}
	book := &Book{}
	if getBookInfo(book, bookId) != nil {
		logger.Log.Error(err)
		return false
	}

	user := &User{}
	if getUserInfo(user, UserId) != nil {
		logger.Log.Error(err)
		return false
	}

	ret.LendTime = time.Now()
	ret.BookId = bookId
	ret.BookName = book.Name
	ret.UserId = UserId
	ret.UserName = user.Name
	err = tx.Create(ret).Error
	if err != nil {
		tx.Rollback()
		logger.Log.Error(err)
		return false
	}
	tx.Commit()
	return true
}

func bookIsLend(bookId int64, is bool) bool {
	if GlobalConn.Table("books").Where("id=?", bookId).Where("lend_out=?", is).First(&Book{}).RowsAffected >= 1 {
		return true
	}
	return false
}

func getBookInfo(ret *Book, id int64) error {
	return GlobalConn.Table("books").Where("id=?", id).First(&ret).Error
}

func getUserInfo(ret *User, id int64) error {
	return GlobalConn.Table("users").Where("id=?", id).First(&ret).Error
}

func UserGiveBook(bookId, userId int64) bool {
	var err error
	tx := GlobalConn.Begin()
	if bookIsLend(bookId, true) == false {
		logger.Log.Error("该书未被借出，归还失败")
		return false
	}

	if bookIsGive(bookId, userId) == true {
		return false
	}

	err = tx.Table("books").Where("id=?", bookId).
		Update("lend_out", false).Error
	if err != nil {
		logger.Log.Error(err)
		tx.Rollback()
		return false
	}

	ti := time.Now()
	err = tx.Table("lend_books").Where("book_id=?", bookId).
		Where("user_id=?", userId).Where("give_book=?", false).
		Update(map[string]interface{}{"give_book": true, "give_time": ti}).Error
	if err != nil {
		tx.Rollback()
		logger.Log.Error(err)
		return false
	}

	tx.Commit()
	return true
}

func bookIsGive(bookId, userId int64) bool {
	if GlobalConn.Table("lend_books").Where("book_id=?", bookId).Where("user_id=?", userId).Where("give_book=?", false).Find(&LendBooks{}).RowsAffected >= 1 {
		return false
	}
	logger.Log.Error("该书已经归还，请勿重复归还")
	return true
}

func GetAllLendInfo(ret *[]APILendBooks, id int64) error {
	return GlobalConn.Table("lend_books").Where("user_id=?", id).Find(&ret).Error
}
