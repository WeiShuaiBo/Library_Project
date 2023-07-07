package modle

import "time"

type User struct {
	UserId   int64  `gorm:"user_id"`
	Name     string `gorm:"name" json:"name" form:"name"`
	Password string `gorm:"password" json:"password" form:"password"`
	Phone    string `gorm:"phone" json:"phone" form:"phone"`
	Tag      int    `gorm:"tag"`
}
type Book struct {
	BookID             int64     `gorm:"primary_key;auto_increment;comment:'书的id'"`
	BookName           string    `gorm:"type:varchar(200);comment:'书名'""`
	Author             string    `gorm:"type:varchar(50);comment:'作者'"`
	PublishingHouse    string    `gorm:"type:varchar(50);comment:'出版社'"`
	Translator         string    `gorm:"type:varchar(50);comment:'译者'"`
	PublishDate        time.Time `gorm:"type:date;comment:'出版时间'"`
	Pages              int       `gorm:"default:100;comment:'页数'"`
	ISBN               string    `gorm:"type:varchar(20);comment:'ISBN号码'"`
	Price              float64   `gorm:"default:1;comment:'价格'"`
	BriefIntroduction  string    `gorm:"type:varchar(15000);default:'';comment:'内容简介'"`
	AuthorIntroduction string    `gorm:"type:varchar(5000);default:'';comment:'作者简介'"`
	ImgURL             string    `gorm:"type:varchar(200);comment:'封面地址'"`
	DelFlg             int       `gorm:"default:0;comment:'删除标识'"`
}
type Borrows struct {
	BorrowId   int64 `gorm:"borrow_id"`
	UserId     int64
	BookId     int64
	BorrowDate time.Time
	ReturnDate time.Time
}
