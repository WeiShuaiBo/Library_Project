package model

import (
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func GetUser(name, pwd string) *User {
	user := &User{}
	sql := "SELECT `id`,`name` from `user` where `UserName` = ? and `password` =? limit 1"
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
		StartTime := time.Now()
		sql = "update book_info set count = count-1 where bookId = ? and count > 0 for update"
		err = tx.Exec(sql, val).Error
		sql = "insert into record (`UserId`,`BookId`,`Status`,`StartTime`,`OverTime`,`ReturnTime`) values(?,?,?,?,?,?)"

		err = tx.Exec(sql, userId, val, state, StartTime, time.Now(), time.Now()).Error
		//GlobalConn.Create(recordNew)

	}
	tx.Exec("commit")

	return true
}
func ReturnBook(userId int64, bookIds []int64) bool {
	err := GlobalConn.Transaction(func(tx *gorm.DB) error {
		for _, val := range bookIds {
			record := &Record{}
			sql := "SELECT * FROM record WHERE BookId = ? and UserId = ?"
			if err := tx.Raw(sql, val, userId).Scan(record).Error; err != nil || record.Status == 0 {
				fmt.Printf("没有找到待还的bookId")
				return err
			}
			sql = "UPDATE record SET Status = 0 WHERE BookId = ? LIMIT 1"
			if err := tx.Exec(sql, val).Error; err != nil {
				return err
			}

		}
		return nil
	})
	if err != nil {
		return false
	}
	return true

}

func Registered(UserName, Password, Name, Sex, Phone string) int {
	userOld := &User{}
	sql := "SELECT * FROM user WHERE UserName = ?"
	err := GlobalConn.Raw(sql, UserName).Scan(userOld).Error
	if err != nil {
		fmt.Printf("查询数据库失败")
		return 1
	}
	if userOld.Id > 0 {
		return 2
	}
	//var worker *Worker
	worker := NewWorker(001, 002)
	//ID:=gg.NextID()
	newId, _ := worker.NextID() // 使用雪花算法生成新的Id
	fmt.Println("newId:")
	fmt.Println(newId)

	sql = "INSERT INTO user (`Id`, `UserName`, `Password`, `Name`, `Sex`, `Phone`, `Status`) VALUES (?, ?, ?, ?, ?, ?, ?)"
	err = GlobalConn.Exec(sql, newId, UserName, Password, Name, Sex, Phone, 0).Error
	if err != nil {
		fmt.Printf("插入数据失败")
		return 1
	}

	return 0
}

func UpdatePersonalInformation(UserName, Password, Name, Sex, Phone string) int {
	userOld := &User{}
	sql := "SELECT * FROM user WHERE UserName = ?"
	err := GlobalConn.Raw(sql, UserName).Scan(userOld).Error
	if err != nil {

		fmt.Printf("查询数据库失败")
		return 1
	}
	if userOld.Id <= 0 {
		return 2
	}
	sql = "UPDATE user SET Password=?, Name=?, Sex=?, Phone=? WHERE UserName=? LIMIT 1"
	if err = GlobalConn.Exec(sql, Password, Name, Sex, Phone, UserName).Error; err != nil {
		return 1
	}
	userIdStr := strconv.FormatInt(userOld.Id, 10)
	_, _ = redisClient.HDel("GetUserInformationMap", userIdStr).Result()
	return 0
}
