// @Author	zhangjiaozhu 2023/7/3 20:40:00
package md5

import (
	"crypto/md5"
	"encoding/hex"
)

var Secret = "md5"

func EncryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(Secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
