package model

import "time"

type User struct {
	Id          int64 `gorm:"id"`
	Name        string
	Pwd         string
	Tel         string
	CreatedTime time.Time
}

type Weather struct {
	Id       int64 `gorm:"id"`
	Time     []string
	Temp     []string
	Humidity []string
	Cloud    []string
}
