package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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
	sql := "SELECT `id`,`UserName` from `librarian` where `UserName` = ? and `Password`=? limit 1"
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

func GetUserInformationRedis(userId int64) *User {
	userIdStr := strconv.FormatInt(userId, 10)
	var user *User
	userData, err := redisClient.HGet(context.Background(), "GetUserInformationMap", userIdStr).Result()
	if err != nil {
		userMysql := GetUserInformation(userId)
		userJson, err := json.Marshal(userMysql)
		if err != nil {
			fmt.Println("转化失败")
			//Response.Error(c,"格式化失败")
			return nil
		}
		err = redisClient.HSet(context.Background(), "GetUserInformationMap", userIdStr, userJson).Err()
		if err != nil {
			//Response.Error(c,"存入redis失败")
			fmt.Printf("存储失败")
			return nil
		}
		return userMysql
	} else {
		err = json.Unmarshal([]byte(userData), &user)
	}
	return user
}

func GetUserInformation(userId int64) *User {
	user := &User{}
	sql := "SELECT * from `user` where `Id` = ? limit 1"
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

//func DeleteBook(bookId int64) int {
//	bookOld := &BookInfo{}
//	sql := "SELECT * FROM book_info WHERE  Id= ?"
//	err := GlobalConn.Raw(sql, bookId).Scan(bookOld).Error
//	if err != nil {
//
//		fmt.Printf("查询数据库失败")
//		return 1
//	}
//	if bookOld.Id > 0 {
//		sql := "DELETE FROM book_info WHERE Id = ?"
//		err := GlobalConn.Exec(sql, bookId).Error
//		if err != nil {
//			// 处理错误
//			return 1
//		}
//		// 删除成功
//		return 0
//	}
//	return 2
//}

// 伪删除
func DeleteBook(bookId int64) int {
	bookOld := &BookInfo{}
	sql := "SELECT * FROM book_info WHERE  Id= ?"
	err := GlobalConn.Raw(sql, bookId).Scan(bookOld).Error
	if err != nil {

		fmt.Printf("查询数据库失败")
		return 1
	}
	if bookOld.Id > 0 {
		sql := "UPDATE book_info SET del_flg = 1 WHERE Id = ?"
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
func DeleteBookRedis(bookId int64) int {
	bookIdString := strconv.FormatInt(bookId, 10)
	// 1.4 数据库存在就存入redis,然后返回前端
	bookNameString, err := redisClient.HGet(context.Background(), "BookIdMap", bookIdString).Result()
	if err != nil {
		fmt.Printf("查询bookNameString失败")
	}
	_, _ = redisClient.LRem(context.Background(), "BookNameList", 0, bookNameString).Result()
	_, _ = redisClient.HDel(context.Background(), "books", bookNameString).Result()
	_, _ = redisClient.HDel(context.Background(), "BookIdMap", bookIdString).Result()

	fmt.Println(bookIdString)
	return 0
}

//
//bookJson, err := json.Marshal()
//if err != nil {
//Response.Error(c,"格式化失败")
//return
//}
//err = redis2.RDB.HSet(c, "books", name.BookName, bookJson).Err()
//if err != nil {
//Response.Error(c,"存入redis失败")
//return
//}
//Response.Success(c,"success",name)
//return 0
