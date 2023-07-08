package model

import (
	"library/appV0/logger"
	"time"
)

func GetUser(name, pwd string) *User {
	ret := &User{}
	sql := "SELECT * FROM user WHERE name=?,pwd=?"
	if err := GlobalConn.Raw(sql, name, pwd).Scan(&ret).Error; err != nil {

	}
	if err := GlobalConn.Table("users").Where("name = ? and pwd = ?", name, pwd).First(ret).Error; err != nil {
		logger.Log.Error("查询错误", err.Error())
	}
	return ret
}

func GetAdmin(name, pwd string) *User {
	ret := &User{}
	sql := "SELECT * FROM admin WHERE name=?,pwd=?"
	if err := GlobalConn.Raw(sql, name, pwd).Scan(&ret).Error; err != nil {
		logger.Log.Error(err)
	}
	return ret
}

func CheckGetUserExist(name string) bool {
	ret := &User{}
	sql := "SELECT * FROM user WHERE user_name=?"
	if GlobalConn.Raw(sql, name).Scan(&ret).RowsAffected <= 0 {
		return false
	}
	return true
}

func RegisterUser(u *User) error {
	sql := "INSERT INTO user (name,pwd,tel) values (?,?,?)"
	if err := GlobalConn.Exec(sql, u.Name, u.Pwd, u.Tel).Error; err != nil {
		return err
	}
	return nil
}

func GetInfo(ret *APIUser, id int64) error {
	sql := "SELECT * FROM user WHERE id=?"
	if err := GlobalConn.Raw(sql, id).Scan(&ret).Error; err != nil {
		return err
	}
	return nil
}

func UpdateInfo(ret *APIUser, id int64) error {
	sql := "UPDATE user set name=?,pwd=?,tel=? WHERE id=?"
	if err := GlobalConn.Exec(sql, ret.Name, ret.Pwd, ret.Tel, id).Error; err != nil {
		return err
	}
	return nil
}

//借书

func UserLendBook(bookId, UserId int64) bool {
	var err error
	tx := GlobalConn.Begin()
	if bookIsLend(bookId, true) == true {
		logger.Log.Error("该书已经被借出err:")
		return false
	}
	sql := "UPDATE book_info SET lend_out=?,user_id=? WHERE id=?"
	err = tx.Exec(sql, true, UserId, bookId).Error
	if err != nil {
		logger.Log.Error(err)
		tx.Rollback()
		return false
	}

	ret := &UserBook{}
	book := &BookInfo{}
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
	ret.BookName = book.BookName
	ret.UserId = UserId
	ret.UserName = user.Name
	sql = "INSERT INTO user_book (book_id,book_name,user_id,user_name,lend_time) VALUES (?,?,?,?,?)"
	err = tx.Exec(sql, ret.BookId, ret.BookName, ret.UserId, ret.UserName, ret.LendTime).Error
	if err != nil {
		tx.Rollback()
		logger.Log.Error(err)
		return false
	}
	tx.Commit()
	return true
}

func bookIsLend(bookId int64, is bool) bool {
	ret := &UserBook{}
	sql := "SELECT * FROM book_info WHERE id=?,lend_out=?"
	if GlobalConn.Raw(sql, bookId, is).Scan(&ret).RowsAffected >= 1 {
		return true
	}
	return false
}

func getBookInfo(ret *BookInfo, id int64) error {
	sql := "SELECT * FROM book_info WHERE id=?"
	if err := GlobalConn.Raw(sql, id).Scan(&ret).Error; err != nil {
		return err
	}
	return nil
}

func getUserInfo(ret *User, id int64) error {
	sql := "SELECT * FROM user WHERE id=?"
	if err := GlobalConn.Raw(sql, id).Scan(&ret).Error; err != nil {
		return err
	}
	return nil
}

//还书

func UserGiveBook(bookId, userId int64) bool {
	var err error
	tx := GlobalConn.Begin()
	if bookIsLend(bookId, true) == false {
		logger.Log.Error("该书未被借出，归还失败")
		return false
	}

	if bookIsGive(bookId, userId) == true {
		logger.Log.Error("该书已经归还，请勿重复归还")
		return false
	}

	sql := "UPDATE book_info SET lend_out=?,user_id=? WHERE id=?"
	err = tx.Exec(sql, false, "", bookId).Error
	if err != nil {
		logger.Log.Error(err)
		tx.Rollback()
		return false
	}

	ti := time.Now()
	sql = "UPDATE user_book SET give_book=?,give_time=? WHERE book_id=?,user_id=?,give_book=?"
	err = tx.Exec(sql, true, ti, bookId, userId, false).Error
	if err != nil {
		tx.Rollback()
		logger.Log.Error(err)
		return false
	}

	tx.Commit()
	return true
}

func bookIsGive(bookId, userId int64) bool {
	ret := &UserBook{}
	sql := "SELECT * FROM user_book WHERE book_id=?,user_id=?,give_book=?"
	if GlobalConn.Raw(sql, bookId, userId, false).Scan(&ret).RowsAffected >= 1 {
		return false
	}
	return true
}

func GetAllLendInfo(ret *[]APIUserBook, id int64) error {
	sql := "SELECT * FROM user_book WHERE id=?"
	if err := GlobalConn.Raw(sql, id).Scan(&ret).Error; err != nil {
		return err
	}
	return nil
}
