// @Author	zhangjiaozhu 2023/7/5 9:11:00
package dao

import (
	"MyBook/models"
	"MyBook/models/DB/MysqlDB"
	"errors"
	"gorm.io/gorm"
)

func FindUserRecord(id uint64) ([]*models.BorrowRecord, error) {
	record := make([]*models.BorrowRecord, 10)
	err := MysqlDB.DB.Where("user_id = ?", id).Find(&record).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, errors.New("服务器繁忙")
		}
	}
	if len(record) == 0 {
		return nil, errors.New("查无此记录")
	}
	return record, nil
}

func GetAllRecord() ([]*models.BorrowRecord, error) {
	records := make([]*models.BorrowRecord, 10)
	err := MysqlDB.DB.Find(&records).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("服务器繁忙")
	}
	if len(records) == 0 {
		return nil, errors.New("空记录")
	}
	return records, nil
}

func GetRecordByUser(s string) ([]*models.BorrowRecord, error) {
	record := make([]*models.BorrowRecord, 10)
	err := MysqlDB.DB.Where("user_name = ? ", s).Find(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("服务器繁忙")
	}
	if len(record) == 0 {
		return nil, errors.New("空记录")
	}
	return record, nil
}
