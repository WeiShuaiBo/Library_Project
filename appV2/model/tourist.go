package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// 添加BookAllSet
//func GetAllRedis2(pageNumber int) []*BookInfo {
//	fmt.Println(pageNumber)
//	bookInfo := make([]*BookInfo, 0)
//	sql := "select * from book_Info"
//	err := GlobalConn.Raw(sql).Scan(&bookInfo).Error
//	if err != nil {
//		fmt.Printf("查询全部失败")
//	}
//	fmt.Println(len(bookInfo))
//	fmt.Printf("bookInfo[1].Id：%d，类型：%T\n", bookInfo[1].Id, bookInfo[1].Id)
//	for _, book := range bookInfo {
//		redisClient.RPush("BookNameMap")
//
//		//z := redis.Z{Score: float64(book.Id), Member: book.BookName}
//		//redisClient.ZAdd("BookAllSet", z)
//	}
//	return nil
//
//}

// 分页查询的getAll
func GetAll(start1 int) []*BookInfo {
	fmt.Printf("GetAll进入成功")
	bookInfo := make([]*BookInfo, 0)
	//pageSize := 10 //每页显示的数据条数。
	pageSize := 10
	//pageNumber := pageNumber                       //要查询的页码。
	//offset := (pageNumber - 1) * pageSize //偏移量，用于确定从数据库中的哪一行开始获取数据。

	var total int
	result := GlobalConn.Raw("SELECT COUNT(*) FROM book_Info").Scan(&total)
	if result.Error != nil {
		fmt.Println("Failed to count books:", result.Error)
		return bookInfo
	}

	if start1 >= total {
		fmt.Println("Invalid page number")
		return bookInfo
	}

	//sql := "select * from book_Info LIMIT ? OFFSET ?"
	sql := "select * from book_Info where id > ? LIMIT ? "
	err := GlobalConn.Raw(sql, start1, pageSize).Scan(&bookInfo).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		fmt.Println("err！=nil")
		return bookInfo
	}
	return bookInfo
}

func GetBook(id int64) *BookInfo {
	var book *BookInfo

	sql := "select * from book where id = " + strconv.FormatInt(id, 10)

	err := GlobalConn.Raw(sql).Scan(&book).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return book
	}
	return book
}
func GetBookByName(bookName string) *BookInfo {
	var book *BookInfo
	fmt.Printf("GetBookByName进入成功")
	fmt.Print(bookName)
	sql := "select * from book_info where book_name LIKE ?"

	err := GlobalConn.Raw(sql, "%"+bookName+"%").Scan(&book).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return book
	}
	return book
}
func GetBookByNameRedis(bookName string) *BookInfo {
	fmt.Printf("GetBookByNameRedis进入成功")
	var book *BookInfo
	bookData, err := redisClient.HGet("books", bookName).Result()
	fmt.Println("qian")
	fmt.Println(bookData)
	fmt.Println("后")
	if bookData == "" {
		book = GetBookByName(bookName)
		//存储book
		if book != nil {
			bookJSON, err := json.Marshal(book)
			if err != nil {
				fmt.Printf("书籍信息序列化失败: %v\n", err)
				return nil
			}
			err = redisClient.HSet("books", bookName, string(bookJSON)).Err()
			if err != nil {
				fmt.Printf("redis存储书籍信息出错: %v\n", err)
			}
		}

	} else {

		err = json.Unmarshal([]byte(bookData), &book)
		if err != nil {
			fmt.Printf("转换失败")
		}
	}
	return book
}

func QrcodeRedis(code string) bool {
	var User []*User
	userJson, err := json.Marshal(User)
	if err != nil {
		fmt.Printf("转化失败")
	}
	err = redisClient.HSet("QrcodeMap", code, userJson).Err()
	if err != nil {
		//Response.Error(c,"存入redis失败")
		fmt.Printf("存储失败")
		return false
	}
	return true
}

func AutomaticLogin(code string) *User {
	user := &User{}
	userData, err := redisClient.HGet("QrcodeMap", code).Result()
	if err != nil {
		fmt.Printf("获取失败")
		return nil
	}
	err = json.Unmarshal([]byte(userData), user)
	if err != nil {
		fmt.Printf("转化失败")
		return nil
	}
	return user
}
