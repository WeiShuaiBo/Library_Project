package model

type Book struct {
	ID         uint64 `json:"id" gorm:"id"`
	BookName   string `json:"bookname" gorm:"book_name"`
	BookAuthor string `json:"bookauthor" gorm:"book_author"`
	BookNumber int    `json:"booknumber" gorm:"book_number"`
	BookKind   string `json:"bookkind" gorm:"book_kind"`
	BookBrief  string `json:"bookbrief" gorm:"book_brieg"`
}
