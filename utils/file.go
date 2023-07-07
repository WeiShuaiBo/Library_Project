package utils

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
)

var (
	basePath   string
	BatPath    string //bat脚本存储位置
	ShellPath  string //shell脚本存储位置
	PythonPath string //python脚本存储位置
)

//初始化脚本任务存储路径

func GetProjectPath() {
	_, filename, _, ok := runtime.Caller(0)

	if ok {
		basePath = path.Dir(filename)
	}
	project := strings.Split(basePath, "/")
	basePath = strings.Join(project[:len(project)-1], "/")
	BatPath = basePath + "/scripts/bat/"
	ShellPath = basePath + "/scripts/shell/"
	PythonPath = basePath + "/scripts/python/"
}

// 将接受到的数据写入文件
func WriteFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(data)
	if err != nil {
		return err
	}
	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	return err
}

// ToDto 将数据从 from 拷贝到 to
// binding type interface 要修改的结构体
// value type interace 有数据的结构体
// author ziop(http://github.com/wangpi26)
func ToDto(from interface{}, to interface{}) {
	toval := reflect.ValueOf(to).Elem()     //获取reflect.Type类型
	fromval := reflect.ValueOf(from).Elem() //获取reflect.Type类型
	vTypeOfT := fromval.Type()
	for i := 0; i < fromval.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := toval.FieldByName(name).IsValid(); ok {
			toval.FieldByName(name).Set(reflect.ValueOf(fromval.Field(i).Interface()))
		}
	}
}
