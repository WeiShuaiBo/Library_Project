package model

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"sync"
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

		sql = "insert into record (`UserId`,`BookId`,`Status`,`StartTime`,`OverTime`,`ReturnTime`) values(?,?,?,?,?,?)"

		err = tx.Exec(sql, userId, val, state, StartTime, time.Now(), time.Now()).Error
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
	_, _ = redisClient.HDel(context.Background(), "GetUserInformationMap", userIdStr).Result()
	return 0
}

func BuyBook(bookId, userId int64) *Order {
	fmt.Printf("成共进入到buybook")
	var rateLock sync.Mutex
	rateLock.Lock()
	defer rateLock.Unlock()
	tx := GlobalConn.Begin()

	couriers := make([]*Courier, 0)
	sql := "SELECT * FROM courier WHERE status = 1"
	err := GlobalConn.Raw(sql).Scan(&couriers).Error
	if err != nil {
		fmt.Printf("获取courierId失败")
		return nil
	}
	fmt.Printf("数据库查询成果")
	// 设置随机种子，确保不同的运行生成不同的随机数序列
	rand.Seed(time.Now().UnixNano())

	// 生成一个随机整数，范围在 0 到 100 之间
	randomNumber := rand.Intn(len(couriers))
	fmt.Println(randomNumber)
	worker := NewWorker(001, 002)
	//ID:=gg.NextID()
	orderId, _ := worker.NextID() // 使用雪花算法生成新的Id
	fmt.Println("orderId:")
	fmt.Println(orderId)
	sql = "INSERT INTO orders (`order_id`, `book_id`, `courier_id`, `order_status`, `created_at`,`user_id`) VALUES (?, ?, ?, ?, ?,?)"
	err = GlobalConn.Exec(sql, orderId, bookId, couriers[randomNumber].Id, "待支付", time.Now(), userId).Error
	if err != nil {
		fmt.Printf("生成订单失败")
		return nil
	}
	sql = "UPDATE book_info SET count = count-1 WHERE count > 0 and id= ? LIMIT 1"
	if err = GlobalConn.Exec(sql, bookId).Error; err != nil {
		return nil
	}
	orders := &Order{}
	sql = "SELECT * FROM orders WHERE order_id = ?"
	err = GlobalConn.Raw(sql, orderId).Scan(&orders).Error
	if err != nil {
		fmt.Printf("获取courierId失败")
		return nil
	}
	tx.Exec("commit")
	return orders
}

func GetBuyBook() []*Order {
	orders := make([]*Order, 0)
	//sql := "SELECT * FROM orders WHERE delete = ?"
	sql := "SELECT * FROM orders WHERE del_flg = 0"
	err := GlobalConn.Raw(sql).Scan(&orders).Error
	if err != nil {

		return nil
	}
	return orders
}
func NoticePay(userId int64) *Order {
	var rateLock sync.Mutex
	rateLock.Lock()
	defer rateLock.Unlock()
	tx := GlobalConn.Begin()
	order := &Order{}
	sql := "SELECT * FROM orders WHERE user_id = ? and order_status = '未通知尽快支付'"
	err := tx.Raw(sql, userId).Scan(&order).Error
	if err != nil {
		fmt.Printf("获取courierId失败")
		return nil
	}
	if order.OrderId <= 0 {
		return nil
	}
	sql = "UPDATE orders SET order_status = '已通知尽快支付' WHERE user_id = ? and order_status = '未通知尽快支付'"
	if err = GlobalConn.Exec(sql, userId).Error; err != nil {
		return nil
	}
	tx.Exec("commit")
	return order
}
func NoticeOrders(OrderId int64) bool {
	sql := "UPDATE orders SET order_status = '未通知尽快支付' WHERE order_id = ? LIMIT 1"
	if err := GlobalConn.Exec(sql, OrderId).Error; err != nil {
		return false
	}
	return true
}

func CancelBuyBook(id int64) {
	sql := "UPDATE orders SET del_flg= 1 WHERE order_id = ? LIMIT 1"
	if err := GlobalConn.Exec(sql, id).Error; err != nil {
		fmt.Printf("取消订单出现问题")
	}

}

func QrcodeLogin(code string, userId int64, userName string) bool {
	fmt.Printf("model进入成功")
	fmt.Println(code)
	fmt.Println(userId)
	fmt.Println(userName)

	user := &User{}
	user.Id = userId
	fmt.Println(user.Id)
	user.Name = userName
	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println("转化失败")
		return false
	}
	err = redisClient.HSet(context.Background(), "QrcodeMap", code, userJson).Err()
	if err != nil {
		fmt.Printf("更新失败")
	}
	return true
}

func GetPay(outTradeNo string) bool {
	fmt.Printf("进入GetPay")
	var rateLock sync.Mutex
	rateLock.Lock()
	defer rateLock.Unlock()
	tx := GlobalConn.Begin()
	orders := &Order{}
	sql := "SELECT * FROM orders WHERE order_id = ? and del_flg = 0"
	err := tx.Raw(sql, outTradeNo).Scan(&orders).Error
	if err != nil {
		fmt.Printf("查询错误1")
		return false
	}
	if orders.OrderId <= 0 {
		fmt.Printf("查询错误2")
		return false
	}
	fmt.Print(orders.OrderId)
	tx.Exec("commit")
	return true
}

func Pay(outTradeNo string) bool {
	fmt.Printf("进入Pay")

	var rateLock sync.Mutex
	rateLock.Lock()
	defer rateLock.Unlock()
	tx := GlobalConn.Begin()
	sql := "UPDATE orders SET order_status = '代发货' WHERE order_id = ? LIMIT 1"
	if err := tx.Exec(sql, outTradeNo).Error; err != nil {
		return false
	}
	tx.Exec("commit")
	return true
}
