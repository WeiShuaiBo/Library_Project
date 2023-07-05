// @Author	zhangjiaozhu 2023/7/4 10:31:00
package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	BookId    uint64
	BookName  string // 书名
	Author    string // 作者
	Price     string // 价格
	Type      string // 类型
	Reserve   string // 预约数量
	Loan      string // 借出数量
	Remaining string // 剩余本数
	Desc      string // 简介
}

func (book *Book) TableName() string {
	return "book"
}
