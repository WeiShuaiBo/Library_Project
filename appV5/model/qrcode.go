package model

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

func GetQRcode(data string, name string) bool {
	err := qrcode.WriteFile(data, qrcode.Medium, 256, name)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("success")
	return true
}
