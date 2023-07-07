package utils

import (
	"errors"
	"os"
)

func PathExists(path string) (bool, error) {
	// os.Stat 返回文件的FileInfo的值
	fi, err := os.Stat(path)
	if err == nil {
		// IsDir() 判断路径是否表示文件夹。？？
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
