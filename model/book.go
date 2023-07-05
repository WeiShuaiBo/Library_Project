package model

import "fmt"

func GetBooks(ret *[]APIBook) error {
	if err := GlobalConn.Table("books").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s\n", err)
		return err
	}
	return nil
}
func GetBookByKeyWord(ret *[]APIBook, keyWord string) error {
	if err := GlobalConn.Table("books").Where("name like ?", "%"+keyWord+"%").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s\n", err)
		return err
	}
	return nil
}
