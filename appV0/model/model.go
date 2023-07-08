package model

import "time"

type User struct {
	Id          int64     `gorm:"id"`
	Name        string    `form:"name" binding:"required"` //用户名
	Pwd         string    `form:"pwd" binding:"required"`  //密码
	Tel         string    `form:"tel" binding:"required"`  //手机号
	CreatedTime time.Time //注册时间
}

func (User) TableName() string {
	return "user"
}

type Admin struct {
	Id          int64     `gorm:"id"`
	Name        string    `form:"name" binding:"required"` //用户名
	Pwd         string    `form:"pwd" binding:"required"`  //密码
	Tel         string    `form:"tel" binding:"required"`  //手机号
	CreatedTime time.Time //注册时间
}

func (Admin) TableName() string {
	return "admin"
}

//type Book struct {
//	Id      int64  `gorm:"id"`
//	Name    string `json:"name" binding:"required"`   //图书名字
//	Author  string `json:"author" binding:"required"` //图书作者
//	Number  string `json:"number" binding:"required"` //图书编号
//	LendOut bool   `json:"lend_out"`                  //图书是否借出
//	UserId  int64  `json:"user_id"`                   //借出用户
//	//Count   int64  `json:"count"`                   //被借次数
//}

//func (Book) TableName() string {
//	return "book"
//}

type BookInfo struct {
	Id                 int64     `gorm:"id"`                  //书的id
	BookName           string    `json:"book_name"`           //书名
	Author             string    `json:"author"`              //作者
	PublishingHouse    string    `json:"publishing_house"`    //出版社
	Translator         string    `json:"translator"`          //译者
	PublishDate        time.Time `json:"publish_date"`        //出版时间
	Pages              int64     `json:"pages"`               //页数
	ISBN               string    `json:"isbn"`                //ISBN号码
	Price              float64   `json:"price"`               //价格
	BriefIntroduction  string    `json:"brief_introduction"`  //内容简介
	AuthorIntroduction string    `json:"author_introduction"` //作者简介
	ImgUrl             string    `json:"img_url"`             //封面地址
	DelFlg             int64     `json:"del_flg"`             //删除标识
	LendOut            bool      `json:"lend_out"`            //图书是否借出
	UserId             int64     `json:"user_id"`             //借出用户
}

func (BookInfo) TableName() string {
	return "book_info"
}

type UserBook struct {
	Id       int64     `gorm:"id"`
	BookId   int64     //图书Id
	BookName string    //图书名
	UserId   int64     //用户Id
	UserName string    //用户名
	LendTime time.Time //借出时间
	GiveBook bool      //是否归还
	GiveTime time.Time //归还时间
}

func (UserBook) TableName() string {
	return "user_book"
}

type APIUser struct {
	Name string `json:"name" form:"name" binding:"required"`
	Pwd  string `json:"pwd" form:"pwd" binding:"required"`
	Tel  string `json:"tel" form:"tel" binding:"required"`
}

type APIBookInfo struct {
	BookName          string `json:"book_name"`          //书名
	Author            string `json:"author"`             //作者
	PublishingHouse   string `json:"publishing_house"`   //出版社
	BriefIntroduction string `json:"brief_introduction"` //内容简介
	ImgUrl            string `json:"img_url"`            //封面地址
	LendOut           bool   `json:"lend_out"`           //图书是否借出
}

type APIUserBook struct {
	BookName string
	UserName string
	LendTime time.Time
	GiveBook bool
	GiveTime time.Time
}
