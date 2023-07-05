package model

// CreatLibrary 添加图书
func CreatLibrary(library *Library) error {

	//return GlobalConn.Table("books").Create(&library).Error

	sql := "INSERT INTO books (title, author, publisher, edition, stock, user_id) VALUES (?, ?, ?, ?, ?, ?)"

	return GlobalConn.Exec(sql, library.Title, library.Author, library.Publisher, library.Edition, library.Stock, library.UserId).Error
}

// GetLibrary 获取图书
func GetLibrary(page, pageSize int) ([]Library, error) {

	var library []Library
	//sql语句 sql := "SELECT * FORM books LIMIT ? OFFSET ?"
	err := GlobalConn.Table("books").Limit(pageSize).Offset((page - 1) * pageSize).Find(&library).Error
	if err != nil {
		return nil, err
	}

	return library, nil
}

func UpdateLibrary(library *Library) error {
	// 构建 UPDATE SQL 语句
	sql := "UPDATE books SET title = ?, author = ?, publisher = ?, edition = ?, stock = ? WHERE id = ?"

	// 执行 SQL 语句
	return GlobalConn.Table("books").Exec(sql, library.Title, library.Author, library.Publisher, library.Edition, library.Stock, library.Id).Error

}

func DeleteLibrary(id int64) error {

	sql := "DELETE FROM books WHERE id = ?"
	return GlobalConn.Table("books").Exec(sql, id).Error

	//return GlobalConn.Table("books").Delete("id = ?", id).Error
}

// FindLibrary 模糊查询
func FindLibrary(title string) ([]Library, error) {
	var library []Library
	// 构建 SQL 语句
	sql := "SELECT * FROM books WHERE title LIKE ?"
	// 执行 SQL 查询
	if err := GlobalConn.Table("books").Raw(sql, "%"+title+"%").Scan(&library).Error; err != nil {
		return nil, err
	}
	return library, nil
}
