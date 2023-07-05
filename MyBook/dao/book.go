// @Author	zhangjiaozhu 2023/7/4 10:27:00
package dao

import (
	"MyBook/models"
	"MyBook/models/DB/MysqlDB"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func GetBook() ([]*models.Book, error) {
	data := make([]*models.Book, 10)
	err := MysqlDB.DB.Find(&data).Debug().Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
func FindBookByName(s string) (models.Book, error) {
	var book models.Book
	err := MysqlDB.DB.Where("book_name = ?", s).First(&book).Debug().Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return models.Book{}, errors.New("服务器繁忙")
		}
	}
	return book, nil
}

func FindBookById(id uint64) (models.Book, error) {
	var book models.Book
	err := MysqlDB.DB.Where("id=?", id).First(&book).Debug().Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Book{}, errors.New("查无此书")
		}
		return models.Book{}, errors.New("服务器繁忙")
	}
	return book, nil
}

func CreateBook(book *models.Book) error {
	if err := MysqlDB.DB.Create(&book).Debug().Error; err != nil {
		return errors.New("服务器繁忙")
	}
	return nil
}

func DeleteBookByName(book models.Book) error {
	// 逻辑删除
	MysqlDB.DB.Where("book_id = ?", book.BookId).Delete(&book)
	// 检查记录是否被删除
	result := MysqlDB.DB.First(&book)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil
		} else {
			return errors.New("删除记录失败")
		}
	}
	return nil
}

func BorrowBook(id uint64, book *models.Book) error {
	var br models.BorrowRecord
	br = models.BorrowRecord{
		UserId: id,
		BookId: book.BookId,
		Status: 1,
	}
	atoi, err := strconv.Atoi(book.Remaining)
	if err != nil {
		return errors.New("服务器繁忙")
	}
	// 图书库存-1
	book.Remaining = strconv.Itoa(atoi - 1)

	// 查询是否存在借阅记录
	first, err := FindBorrowBookFirst(id, book.BookId)
	if err != nil {
		return err
	}
	if first.ID == 0 {
		MysqlDB.DB.Save(book)
		MysqlDB.DB.Create(&br)
		return nil
	}
	// 判断是否归还图书
	if first.Status == 1 {
		return errors.New("请归还图书后方可再次借书")
	}
	if first.Status == 0 {
		MysqlDB.DB.Save(book)
		MysqlDB.DB.Create(&br)
	}
	return errors.New("单用户无法重复借书")

}

func ReturnBook(id uint64, book *models.Book) error {
	MysqlDB.DB.Where("book_name = ?", book.BookName).First(&book)
	atoi, err := strconv.Atoi(book.Remaining)
	if err != nil {
		return errors.New("服务器繁忙")
	}
	// 图书库存+1
	book.Remaining = strconv.Itoa(atoi + 1)
	// 查询是否存在借阅记录
	first, err := FindBorrowBookFirst(id, book.BookId)
	if err != nil {
		return err
	}
	if first.ID < 0 {
		return errors.New("无借阅记录，无法还书")
	}
	first.Status = 0
	MysqlDB.DB.Save(&book)
	MysqlDB.DB.Save(first)
	return nil
}

func FindBorrowBookFirst(userid, bookid uint64) (*models.BorrowRecord, error) {
	var br models.BorrowRecord
	err := MysqlDB.DB.Where("user_id = ? and book_id = ?", userid, bookid).First(&br).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}
	return &br, nil

}

func UpdateBook(book *models.Book) error {
	err := MysqlDB.DB.Updates(book).Error
	if err != nil {
		return err
	}
	return nil
}
