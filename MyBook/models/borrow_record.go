// @Author	zhangjiaozhu 2023/7/4 10:38:00
package models

import "gorm.io/gorm"

type BorrowRecord struct {
	gorm.Model
	UserId uint64
	BookId uint64
	Status int
}

func (br *BorrowRecord) TableName() string {
	return "borrow_record"
}
