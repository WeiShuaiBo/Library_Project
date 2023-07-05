package model

import (
	"fmt"
)

// FindRecordByUserId 查询自己的最新借书记录
func FindRecordByUserId(userId int64) *Borrow {

	sql := "SELECT * FROM borrow WHERE user_id = ? ORDER BY id DESC LIMIT 1"
	ret := &Borrow{}
	err := GlobalConn.Table("borrow").Exec(sql, userId).First(&ret).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
	}
	return ret
}

// GetRecordByUserId 查询用户借阅记录
func GetRecordByUserId(userId int64) ([]Borrow, error) {
	sql := "SELECT * FROM borrow WHERE user_id = ?"
	var borrow []Borrow
	// 执行 SQL 查询
	if err := GlobalConn.Table("books").Raw(sql, userId).Scan(&borrow).Error; err != nil {
		return nil, err
	}
	return borrow, nil
}

// GetRecord 查询用户借阅记录
func GetRecord() ([]Borrow, error) {
	sql := "SELECT * FROM borrow"
	var borrow []Borrow
	// 执行 SQL 查询
	if err := GlobalConn.Table("books").Raw(sql).Scan(&borrow).Error; err != nil {
		return nil, err
	}
	return borrow, nil
}

// GetExpected 查询预期的用户
func GetExpected() ([]Borrow, error) {

	var borrow []Borrow

	sql := "SELECT * FROM Borrow WHERE is_return= false AND due_date > NOW()"

	if err := GlobalConn.Table("borrow").Raw(sql).Scan(&borrow).Error; err != nil {
		return nil, err
	}
	return borrow, nil
}
