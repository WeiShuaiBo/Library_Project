// @Author	zhangjiaozhu 2023/7/5 9:09:00
package service

import (
	"MyBook/dao"
	"MyBook/models"
)

func FindUserRecord(id uint64) ([]*models.BorrowRecord, error) {
	return dao.FindUserRecord(id)
}
func GetAllRecord() ([]*models.BorrowRecord, error) {
	return dao.GetAllRecord()
}

func GetRecordByUser(s string) ([]*models.BorrowRecord, error) {
	return dao.GetRecordByUser(s)
}
