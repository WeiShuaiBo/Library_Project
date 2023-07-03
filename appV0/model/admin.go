package model

import "fmt"

func GetRecords() []*Record {
	ret := make([]*Record, 0)
	sql := "select * from `Record` where id > 0"
	err := GlobalConn.Raw(sql).Scan(&ret).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return ret
	}
	return ret
}

func GetAdmin(name, pwd string) *Librarian {
	user := &Librarian{}
	sql := "SELECT `id`,`name` from `librarian` where `name` = ? and `Password`=? limit 1"
	err := GlobalConn.Raw(sql, name, pwd).Scan(user).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())

	}
	return user
}
func GetUserRecordStatus(id int64) []*Record {
	ret := make([]*Record, 0)
	sql := "select * from Record where Status = ?"
	err := GlobalConn.Raw(sql, id).Scan(&ret).Error
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return ret
	}
	return ret
}
