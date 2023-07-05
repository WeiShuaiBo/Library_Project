package model

type UserBook struct {
	ID        uint64 `json:"id" gorm:"id"`
	UserId    uint64 `json:"userid" gorm:"user_id"`
	BookId    uint64 `json:"bookid" gorm:"book_id"`
	IsReturn  string `json:"isreturn" gorm:"is_return"`
	StartTime string `json:"starttime" gorm:"start_time"`
	EndTime   string `json:"endtime" gorm:"end_time"`
}
