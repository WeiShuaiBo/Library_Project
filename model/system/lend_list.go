package system

import "time"

type LendList struct {
	Id       int       `json:"Id" gorm:"id"`
	UserId   int       `json:"UserId" gorm:"user_id" `
	BookId   int       `json:"BookId" gorm:"book_id"`
	LendTime time.Time `json:"LendTime" gorm:"lend_time"`
	EndTime  time.Time `json:"EndTime" gorm:"end_time"`
}
