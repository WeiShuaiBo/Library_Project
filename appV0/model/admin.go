package model

import (
	"fmt"
)

func GetRecords() []*Record {
	ret := make([]*Record, 0)
	sql := "select * from `Record` where id > 0"
	err := GlobalConn.Raw(sql).Scan(&ret).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return ret
	}
	return ret
}

func GetAdmin(name, pwd string) *Librarian {
	user := &Librarian{}
	sql := "SELECT `id`,`name` from `librarian` where `name` = ? and `Password`=? limit 1"
	err := GlobalConn.Raw(sql, name, pwd).Scan(user).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())

	}
	return user
}
func GetUserRecordStatus(id int64) []*Record {
	ret := make([]*Record, 0)
	sql := "select * from Record where Status = ?"
	err := GlobalConn.Raw(sql, id).Scan(&ret).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return ret
	}
	return ret
}

func GetUserInformation(userId int64) *User {
	user := &User{}
	sql := "SELECT `id`,`name` from `user` where `Id` = ? limit 1"
	err := GlobalConn.Raw(sql, userId).Scan(user).Error
	if err != nil || user.Id < 0 {
		fmt.Printf("数据库查询有问题")
	}
	return user
}

func AddBook(BN, Name, Description, ImgUrl string, Count int) int {
	bookOld := &Book{}
	sql := "SELECT * FROM book WHERE Name = ?"
	err := GlobalConn.Raw(sql, Name).Scan(bookOld).Error
	if err != nil {

		fmt.Printf("查询数据库失败")
		return 1
	}
	if bookOld.Id > 0 {
		sql = "update book set Count = ? where Name = ? limit 1"
		if err = GlobalConn.Exec(sql, bookOld.Count+Count, Name).Error; err != nil {
			return 1
		}
		return 0
	}
	sql = "insert into book (`BN`,`Name`,`Description`,`ImgUrl`,`Count`,`CategoryId`) values(?,?,?,?,?,?)"

	err = GlobalConn.Exec(sql, BN, Name, Description, ImgUrl, Count, 0).Error
	return 0
}

func DeleteBook(bookId int64) int {
	bookOld := &Book{}
	sql := "SELECT * FROM book WHERE  Id= ?"
	err := GlobalConn.Raw(sql, bookId).Scan(bookOld).Error
	if err != nil {

		fmt.Printf("查询数据库失败")
		return 1
	}
	if bookOld.Id > 0 {
		sql := "DELETE FROM book WHERE Id = ?"
		err := GlobalConn.Exec(sql, bookId).Error
		if err != nil {
			// 处理错误
			return 1
		}
		// 删除成功
		return 0
	}
	return 2
}

func GETBookRecord(bookId int64) *Record {
	ret := &Record{}
	sql := "select * from `record` where BookId = ?"
	err := GlobalConn.Raw(sql, bookId).Scan(&ret).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return ret
	}
	return ret
}
