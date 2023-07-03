package model

import (
	"fmt"
	"time"
)

func GetUser(name, pwd string) *User {
	user := &User{}
	sql := "SELECT `id`,`name` from `user` where `name` = ? and `password` =? limit 1"
	err := GlobalConn.Raw(sql, name, pwd).Scan(user).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())

	}
	return user
}

func BorrowBook(userId int64, bookIds []int64) bool {
	tx := GlobalConn.Begin()
	for i, val := range bookIds {
		recordOld := &Record{}
		sql := "SELECT * FROM record WHERE BookId = ?"
		err := tx.Raw(sql, val).Scan(recordOld).Error
		fmt.Print("recordOls.Status:")
		fmt.Printf(string(recordOld.Status))
		if err != nil || recordOld.Status != 0 {
			//fmt.Printf("查询数据库有问题")
			return false
		}
		fmt.Print("bookId:")
		fmt.Println(val)
		fmt.Println(i)
		fmt.Print("userId:")
		fmt.Println(userId)
		fmt.Print("time.Now():")
		fmt.Println(time.Now())
		state := 1
		fmt.Print("state:")
		fmt.Println(state)

		//
		//recordNew := &Record{}
		//recordNew.UserId = userId
		//recordNew.BookId = val
		//recordNew.Status = 1
		StartTime := time.Now()

		sql = "insert into record (`UserId`,`BookId`,`Status`,`StartTime`,`OverTime`,`ReturnTime`) values(?,?,?,?,?,?)"

		err = tx.Exec(sql, userId, val, state, StartTime, time.Now(), time.Now()).Error
		//GlobalConn.Create(recordNew)

	}
	tx.Exec("commit")

	return true
}
func ReturnBook(userId int64, bookIds []int64) bool {
	tx := GlobalConn.Begin()
	for i, val := range bookIds {
		record := &Record{}
		sql := "SELECT * FROM record WHERE BookId = ?"
		err := tx.Raw(sql, val).Scan(record).Error

		if err != nil || record.Status == 0 {
			fmt.Printf("没有找到代还的bookId")
			return false
		}
		fmt.Print("bookId:")
		fmt.Println(val)
		fmt.Println(i)
		fmt.Print("userId:")
		fmt.Println(userId)
		fmt.Print("time.Now():")
		fmt.Println(time.Now())
		state := 1
		fmt.Print("state:")
		fmt.Println(state)

		//
		//recordNew := &Record{}
		//recordNew.UserId = userId
		//recordNew.BookId = val
		//recordNew.Status = 1
		//StartTime := time.Now()

		sql = "update record set Status = 0 where BookId = ? limit 1"
		if err = tx.Exec(sql, val).Error; err != nil {
			return false

		}
		//GlobalConn.Create(recordNew)

	}
	tx.Exec("commit")

	return true
}
