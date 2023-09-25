package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"

	//"github.com/go-redis/redis"
	"strconv"
)

func CreateRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址和端口号
		Password: "",               // Redis 访问密码，如果没有设置密码，可以留空
		DB:       0,                // Redis 数据库索引，默认为 0
	})
	//通过 redis.NewClient() 函数创建了一个 Redis 客户端连接。需要提供 Redis 服务器的地址和端口号、密码（如果有设置的话）以及数据库索引。
	//Ping() 方法用于测试与 Redis 的连接是否正常
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		// 处理 Redis 连接错误
		panic(err)
	}

	return client
}

func GetAllRedis(pageNumber int) []*BookInfo {
	pageNumberStr := strconv.Itoa(pageNumber)
	books := make([]*BookInfo, 0)
	//var book *BookInfo
	userData, err := redisClient.HGet(context.Background(), "GetAllMap", pageNumberStr).Result()
	if err != nil {
		BookMysql := GetAll(pageNumber)
		bookJson, err := json.Marshal(BookMysql)
		if err != nil {
			fmt.Println("转化失败")
			//Response.Error(c,"格式化失败")
			return nil
		}
		err = redisClient.HSet(context.Background(), "GetAllMap", pageNumberStr, bookJson).Err()
		if err != nil {
			//Response.Error(c,"存入redis失败")
			fmt.Printf("存储失败")
			return nil
		}
		return BookMysql
	} else {
		err = json.Unmarshal([]byte(userData), &books)
	}
	return books
}

//
//// 以list为索引的尝试
//func GetAllRedis(pageNumber int) []*BookInfo {
//	//books := make([]*Book, 0)
//	fmt.Printf("进入GetAllRedis成功")
//	pageSize := 10 //每页显示的数据条数。
//	//start := int64(pageNumber+1)
//	start := int64((pageNumber - 1) * pageSize)
//	end := start + int64(pageSize) - 1
//
//	//val1, err := redisClient.LRange("BookIdList", start, end).Result()
//	BookIds, err := redisClient.LRange("BookIdList", start, end).Result()
//
//	//var book map[string]interface{}
//	var books []*BookInfo
//	for _, val := range BookIds {
//		// 根据图书Id从 Redis 中获取书籍信息
//		bookData, err := redisClient.HGet("books", val).Result()
//		var book *BookInfo
//		if err != nil {
//			id, _ := strconv.ParseInt(val, 10, 64)
//			book = GetBook(id)
//			bookJson, err := json.Marshal(book)
//			if err != nil {
//				fmt.Println("转化失败")
//				//Response.Error(c,"格式化失败")
//				return nil
//			}
//			err = redisClient.HSet("books", string(book.Id), bookJson).Err()
//			if err != nil {
//				//Response.Error(c,"存入redis失败")
//				fmt.Printf("存储失败")
//				return nil
//			}
//			//continue
//		} else {
//			err = json.Unmarshal([]byte(bookData), &book)
//		}
//
//		// 将书籍信息转换为 BookInfo 结构体
//		//book := &BookInfo{
//		//	Id:                 convertToInt(string(bookData["id"])),
//		//	BookName:           bookData["BookName"],
//		//	Author:             bookData["Author"],
//		//	BriefIntroduction:  bookData["Description"],
//		//	Price:              convertToFloat(bookData["Price"]),
//		//	PublishingHouse:    bookData["PublishingHouse"],
//		//	Translator:         bookData["Translator"],
//		//	PublishDate:        convertToTime(bookData["PublishDate"]),
//		//	Pages:              convertToInt(bookData["Pages"]),
//		//	ISBN:               bookData["ISBN"],
//		//	AuthorIntroduction: bookData["AuthorIntroduction"],
//		//	ImgUrl:             bookData["ImgUrl"],
//		//	DelFlg:             convertToInt(bookData["DelFlg"]),
//		//	bookId:             convertToInt64(bookData["bookId"]),
//		//	count:              convertToInt64(bookData["count"]),
//		//}
//
//		books = append(books, book)
//	}
//
//	// 打印书籍信息
//	for _, book := range books {
//		fmt.Printf("书籍信息：%+v\n", book)
//	}
//
//	return books
//
//	fmt.Println("准备查询数据库")
//	bookMySql := GetAll(start, end)
//	if len(bookMySql) == 0 {
//		return nil
//	}
//	//books = bookMySql
//	fmt.Println(bookMySql)
//	redisClient.RPush("BookNameList", bookMySql[0].BookName, bookMySql[1].BookName, bookMySql[2].BookName, bookMySql[3].BookName,
//		bookMySql[4].BookName, bookMySql[5].BookName, bookMySql[6].BookName, bookMySql[7].BookName, bookMySql[8].BookName, bookMySql[9].BookName)
//
//	fmt.Println("准备存储redis")
//	for _, val := range bookMySql {
//		bookIdString := strconv.Itoa(val.Id)
//		err = redisClient.HSet("BookIdMap", bookIdString, val.BookName).Err()
//		if err != nil {
//			fmt.Printf("id索引失败")
//		}
//	}
//
//	err = SaveBooksToRedis(bookMySql)
//	if err != nil {
//		fmt.Println("存储redis失败")
//	}
//	return bookMySql
//}

//func convertToFloat(str string) float64 {
//	if value, err := strconv.ParseFloat(str, 64); err == nil {
//		return value
//	}
//	return 0.0 // 默认值，根据实际情况修改
//}
//func convertToInt(str string) int {
//	if value, err := strconv.Atoi(str); err == nil {
//		return value
//	}
//	return 0 // 默认值，根据实际情况修改
//}
//func convertToInt64(str string) int64 {
//	value, _ := strconv.ParseInt(str, 10, 64)
//	return value
//}
//
//func convertToTime(str string) time.Time {
//	t, _ := time.Parse(time.RFC3339, str)
//	return t
//}

// 根据bookName存储book
func SaveBooksToRedis(books []*BookInfo) error {
	fmt.Println("SaveBooksToRedis进入成功")
	pipe := redisClient.Pipeline()
	defer pipe.Close()

	key := "books"
	existingData, err := redisClient.HGetAll(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	// 创建一个 Map 用于存储书籍信息
	bookMap := make(map[string]interface{})

	for _, book := range books {
		jsonData, err := json.Marshal(book)
		if err != nil {
			return err
		}

		bookMap[book.BookName] = string(jsonData)
	}

	// 将现有的数据与新的书籍信息合并到 bookMap 中
	for field, value := range existingData {
		bookMap[field] = value
	}

	// 存储整个书籍信息的哈希键
	pipe.HMSet(context.Background(), key, bookMap)

	_, err = pipe.Exec(context.Background())
	return err
}

//func SaveBooksToRedis(books []*BookInfo) error {
//	fmt.Println("SaveBooksToRedis进入成功")
//	pipe := redisClient.Pipeline()
//	defer pipe.Close()
//
//	for _, book := range books {
//		jsonData, err := json.Marshal(book)
//		if err != nil {
//			return err
//		}
//
//		key := fmt.Sprintf("books:%d", book.Id)
//		pipe.HSet(key, book.BookName, jsonData)
//	}
//
//	_, err := pipe.Exec()
//	return err
//}
//
//func FuzzySearchRedis(keyword string) []*BookInfo {
//	books := make([]*BookInfo, 0)
//
//	// 使用 HScan 命令模糊查询哈希键，并返回匹配的结果集
//	iter := redisClient.HScan("books", 0, keyword+"*", 0).Iterator()
//	for iter.Next() {
//		field := iter.Val()
//
//		// 根据匹配的哈希键从 Redis 中获取书籍信息
//		bookData, err := redisClient.HGet("books", field).Result()
//		if err != nil {
//			fmt.Printf("获取书籍信息出错: %v\n", err)
//			continue
//		}
//
//		// 将书籍信息转换为 BookInfo 结构体
//		var book BookInfo
//		err = json.Unmarshal([]byte(bookData), &book)
//		if err != nil {
//			fmt.Printf("解析书籍信息出错: %v\n", err)
//			continue
//		}
//
//		books = append(books, &book)
//	}
//	if err := iter.Err(); err != nil {
//		fmt.Printf("模糊查询出错: %v\n", err)
//		return nil
//	}
//
//	// 打印匹配的书籍信息
//	for _, book := range books {
//		fmt.Printf("匹配的书籍信息：%+v\n", book)
//	}
//
//	return books
//}
