package system

import (
	"Library_Project/global"
	"Library_Project/model/system"
	"Library_Project/utils"
	"time"
)

type UserService struct {
}

func (s *UserService) ExistId(id int) bool {
	rowsAffected := global.FAST_DB.Where("id = ?", id).First(&system.User{}).RowsAffected
	if rowsAffected == 1 {
		return true
	}
	return false
}

func (s *UserService) InsertUser(id int, pwd string) {
	password := utils.MD5([]byte(pwd))
	pwd = password
	u := &system.User{Id: id, Password: pwd}
	global.FAST_DB.Create(u)
}

func (s *UserService) Login(id int, pwd string) bool {
	rowsAffected := global.FAST_DB.Where("id = ? and password = ?", id, pwd).First(&system.User{}).RowsAffected
	if rowsAffected == 1 {
		return true
	}
	return false
}

func (s *UserService) Information(id int) *system.User {
	user := &system.User{}

	global.FAST_DB.Where("id = ?", id).First(&user)
	return user
}

func (s *UserService) ChangeInfo(id int, one *system.User) {
	u := &system.User{}
	global.FAST_DB.Where("id = ?", id).First(u)
	one.Id = u.Id
	one.Password = u.Password
	global.FAST_DB.Save(&one)
}

func (s *UserService) Borrow(userId int, book *system.Book) int {
	b := &system.Book{}
	global.FAST_DB.Where("id = ?", book.Id).First(b)
	if b.Statue == 1 {
		//已借出
		return 0
	}
	global.FAST_DB.Model(b).Update("statue", 1)
	//global.FAST_DB.Model(b).Updates(&system.Book{Statue:1})

	list := system.LendList{
		UserId:   userId,
		BookId:   int(book.Id),
		LendTime: time.Now(),
		EndTime:  time.Now().Add(30 * 24 * 60 * time.Minute),
	}

	global.FAST_DB.Create(&list)
	return 1
}

func (s *UserService) Return(book *system.Book) int {
	l := &system.LendList{}
	global.FAST_DB.Where("id = ?", book.Id).First(l)

	if l.EndTime.Before(time.Now()) {
		return 0
	}
	b := &system.Book{}
	global.FAST_DB.Where("id = ?", book.Id).First(b)
	if b.Statue == 0 {
		//未借出
		return 2
	}
	global.FAST_DB.Model(b).Update("statue", 0)

	return 1
}
