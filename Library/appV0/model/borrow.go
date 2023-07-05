package model

// LoanLibrary 借书
func LoanLibrary(borrow *Borrow) error {

	sql := "INSERT INTO borrow (user_id,library_id,loan_date,due_date,is_return) VALUES (?, ?, ?,?,?)"

	return GlobalConn.Exec(sql, borrow.UserId, borrow.LibraryId, borrow.LoanDate, borrow.DueDate, borrow.IsReturn).Error

}

// DueLibrary 还书
func DueLibrary(id int64) error {

	sql := "UPDATE borrow SET due_date= now() , is_return=TRUE WHERE id = ?"

	return GlobalConn.Table("borrow").Exec(sql, id).Error

	//GlobalConn.Table("borrow").Update()
}
