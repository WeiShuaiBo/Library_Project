package model

import "fmt"

func AdminGetBooks(ret *[]Book) error {
	if err := GlobalConn.Table("books").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s\n", err)
		return err
	}
	return nil
}

func AdminGetBookByKeyWord(ret *[]Book, keyWord string) error {
	if err := GlobalConn.Table("books").Where("name like ?", "%"+keyWord+"%").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s\n", err)
		return err
	}
	return nil
}

func AdminGetBooksById(ret *Book, id int64) error {
	if err := GlobalConn.Table("books").Where("id=?", id).First(&ret).Error; err != nil {
		fmt.Printf("err:%s\n", err)
		return err
	}
	return nil
}

func CreatBook(book *Book) error {
	return GlobalConn.Table("books").Create(&book).Error
}

func UpdateBook(book *Book) error {
	if GlobalConn.Table("books").Where("id=?", book.Id).First(&Book{}).RowsAffected <= 0 {
		fmt.Printf("RowsAffected <= 0,更新的数据有误\n")
	}
	return GlobalConn.Table("books").Where("id=?", book.Id).Update(&book).Error
}

func DeleteBook(book *Book, id int64) error {
	if err := GlobalConn.Table("books").Where("id=?", id).First(&book).Error; err != nil {
		fmt.Printf("err:%s\n", err)
		return err
	}
	return GlobalConn.Table("books").Where("id=?", id).Delete(&book).Error
}

func AdminGetInfo(ret *[]LendBooks) error {
	return GlobalConn.Table("lend_books").Find(&ret).Error
}
