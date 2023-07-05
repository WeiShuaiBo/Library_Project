package model

type User struct {
	ID       uint64 `json:"id" gorm:"id"`
	UserName string `json:"username" gorm:"username"`
	PassWord string `json:"password" gorm:"password"`
	Email    string `json:"email" gorm:"email"`
	Phone    uint64 `json:"phone" gorm:"phone"`
	Age      int    `json:"age" gorm:"age"`
	Sex      string `json:"sex" gorm:"sex"`
	Identity string `json:"identity" gorm:"identity"`
}
