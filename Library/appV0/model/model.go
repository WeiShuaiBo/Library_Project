package model

import "time"

type User struct {
	Id          int64  `gorm:"id" json:"id" form:"name"`
	Name        string `json:"name" form:"name" `
	Pwd         string `json:"pwd" form:"pwd" `
	Tel         string `json:"tel" form:"tel"`
	CreatedTime time.Time
	Privilege   int
}

type Library struct {
	Id        int64  `gorm:"id" json:"id" form:"name"`
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
	Edition   string `json:"edition" form:"edition"`
	Stock     int    `json:"stock" form:"stock"`
	UserId    int64  `json:"user_id" form:"user_id"`
}

type Borrow struct {
	Id        int64 `gorm:"id" json:"id" form:"name"`
	UserId    int64 `json:"user_id" form:"user_id"`
	LibraryId int64 `json:"library_id" form:"library_id"`
	LoanDate  time.Time
	DueDate   time.Time
	IsReturn  bool
}

type Page struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
