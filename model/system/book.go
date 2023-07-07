package system

import "time"

type Book struct {
	UUID               uint32     `json:"UUID" gorm:"uuid"`
	Id                 uint32     `json:"Id" gorm:"id"`                                  // 书的id
	BookName           string     `json:"BookName" gorm:"book_name"`                     // 书名
	Author             string     `json:"Author" gorm:"author"`                          // 作者
	PublishingHouse    string     `json:"PublishingHouse" gorm:"publishing_house"`       // 出版社
	Translator         string     `json:"Translator" gorm:"translator"`                  // 译者
	PublishDate        *time.Time `json:"PublishDate" gorm:"publish_date"`               // 出版时间
	Pages              int        `json:"Pages" gorm:"pages"`                            // 页数
	ISBN               string     `json:"ISBN" gorm:"ISBN"`                              // ISBN号码
	Price              float64    `json:"Price" gorm:"price"`                            // 价格
	BriefIntroduction  string     `json:"BriefIntroduction" gorm:"brief_introduction"`   // 内容简介
	AuthorIntroduction string     `json:"AuthorIntroduction" gorm:"author_introduction"` // 作者简介
	ImgUrl             string     `json:"ImgUrl" gorm:"img_url"`                         // 封面地址
	DelFlg             int        `json:"DelFlg" gorm:"del_flg"`                         // 删除标识
	Statue             int        `json:"Statue" gorm:"statue"`
}
