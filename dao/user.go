package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func CheckUser(name, pwd string) User {
	var user User
	sql := "select * from user where name = ? and pwd = ? limit 1"
	err := GlobalConn.Raw(sql, name, pwd).Find(&user).Error
	if err != nil {
		fmt.Println("未找到登录用户信息")
	}
	return user
}
func Register(user User) bool {
	sql := "select * from user where name = ? limit 1"
	var oldUser User
	err := GlobalConn.Raw(sql, user.Name).Find(&oldUser).Error
	if err == nil && user.Name == oldUser.Name {
		fmt.Println("用户信息冲突，请重试")
		return false
	}
	err = GlobalConn.Debug().Table("user").Create(&user).Error
	if err != nil {
		return false
	}
	return true
}
func GetUser(id int) User {
	sql := "select * from user where id = ? limit 1"
	var user User
	if err := GlobalConn.Debug().Raw(sql, id).Find(&user).Error; err != nil {
		fmt.Println("用户信息查询失败")
	}
	return user
}
func PutUser(user User) error {
	sql := "select * from user where name = ?"
	var oldUser User
	err := GlobalConn.Raw(sql, user.Name).Find(oldUser).Error
	if err == nil && oldUser.Name == user.Name {
		fmt.Println("用户名已被注册，请重试")
		return errors.New("用户名冲突")
	}
	return GlobalConn.Where("id = ?", user.Id).Updates(&user).Error
}
func GetMyBook(id int64, page int64) []UserBook {
	var recode []UserBook
	sql := "select * from user_book where user_id = ? limit 10 offset ?"
	if err := GlobalConn.Raw(sql, id, (page-1)*10).Find(&recode).Error; err != nil {
		fmt.Println("查找个人借阅信息失败")
	}
	return recode
}

func Borrow(id int64, bookId int64) bool {

	var book Book
	sql := "select * from book where id = ? limit 1"
	if err := GlobalConn.Raw(sql, bookId).Find(&book).Error; err != nil {
		fmt.Println("未找到指定图书")
		return false
	}
	if book.DelFlg == 1 || book.Count == 0 {
		fmt.Println("指定图书不能借阅")
		return false
	}
	err := GlobalConn.Transaction(func(tx *gorm.DB) error {
		sql1 := "update book set count = count - 1 where id = ?"
		err := tx.Debug().Exec(sql1, bookId).Error
		if err != nil {
			fmt.Println("图书数量减少失败")
			return err
		}
		var newUserBook = UserBook{
			UserId:     id,
			BookId:     book.Id,
			Type:       0,
			BorrowTime: time.Now(),
		}
		if err = tx.Table("user_book").Create(&newUserBook).Error; err != nil {
			fmt.Println("借阅图书操作失败")
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("图书借阅失败")
		return false
	}
	return true
}

func GiveBack(id, bookId int64) bool {
	var book Book
	sql := "select * from book where id = ? limit 1"
	if err := GlobalConn.Raw(sql, bookId).Find(&book).Error; err != nil {
		fmt.Println("未找到指定图书")
		return false
	}
	var userBook UserBook
	sql1 := "select * from user_book where book_id = ? and user_id = ? limit 1"
	if err := GlobalConn.Raw(sql1, bookId, id).Find(&userBook).Error; err != nil {
		fmt.Println("未找到指定图书的借阅记录")
		return false
	}
	err := GlobalConn.Transaction(func(tx *gorm.DB) error {
		sql := "update book set count = count + 1 where id = ?"
		err := tx.Exec(sql, book.Id).Error
		if err != nil {
			fmt.Println("图书数量添加失败")
			return err
		}
		//m := map[string]interface{}{
		//	"type":        1,
		//	"return_time": time.Now(),
		//}
		//if err := GlobalConn.Table("user_book").Where("book_id = ? and user_id = ?", book.Id, id).Updates(m).Error; err != nil {
		//	fmt.Println("归还图书操作失败")
		//	return false
		//}
		sql1 := "update user_book set type = 1,return_time = ? where book_id = ? and user_id = ?"
		err = tx.Exec(sql1, time.Now(), bookId, id).Error
		if err != nil {
			fmt.Println("图书归还记录修改失败")
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return true
}
