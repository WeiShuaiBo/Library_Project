package system

import (
	"Library_Project/global"
	"Library_Project/model/system"
	"Library_Project/model/system/request"
)

type BookService struct {
}

func (s *BookService) Books(p, n int) []*system.Book {
	books := []*system.Book{}
	global.FAST_DB.Offset((p - 1) * n).Limit(n).Find(&books)
	return books
}

func (s *BookService) SearchBooks(str string) []*system.Book {
	books := []*system.Book{}
	global.FAST_DB.Where("book_name LIKE ?", "%"+str).Find(&books)
	return books
}

func (s *BookService) RecordPer(id int) []*request.MyLendBook {
	mybooks := []*request.MyLendBook{}
	//book := []*system.Book{}
	lend := []*system.LendList{}
	global.FAST_DB.Where("user_id = ?", id).Find(&lend)
	for _, list := range lend {
		b := &system.Book{}
		global.FAST_DB.Where("id = ?", list.BookId).First(b)
		temp := &request.MyLendBook{
			*b,
			*list,
		}
		mybooks = append(mybooks, temp)
	}

	return mybooks
}

func (s *BookService) RecordBook(id int) []*request.MyLendBook {
	mybooks := []*request.MyLendBook{}
	lend := []*system.LendList{}
	global.FAST_DB.Where("book_id = ?", id).Find(&lend)
	for _, list := range lend {
		b := &system.Book{}
		global.FAST_DB.Where("id = ?", list.BookId).First(&b)
		temp := &request.MyLendBook{
			*b,
			*list,
		}
		mybooks = append(mybooks, temp)
	}

	return mybooks
}
