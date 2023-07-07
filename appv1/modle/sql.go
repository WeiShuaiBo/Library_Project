package modle

import "fmt"

func GetAllBooks(page int64) ([]Book, error) {
	var books []Book

	offset := (page - 1) * 10 // Calculate offset based on the page number

	sql := "SELECT * FROM books LIMIT 10 OFFSET ?"
	if err := GlobalConn.Raw(sql, offset).Find(&books).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return books, nil
}
func GetABook(book_name string) (Book, error) {
	var books Book
	sql := "SELECT * FROM books WHERE book_name = ? LIMIT 1;"
	if err := GlobalConn.Raw(sql, book_name).First(&books).Error; err != nil {
		fmt.Println(err)
		return books, err
	}
	return books, nil
}
func GetUserInfo(userId any) (User, error) {
	var user User
	sql := "SELECT * FROM users WHERE user_id = ? LIMIT 1"
	if err := GlobalConn.Raw(sql, userId).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
