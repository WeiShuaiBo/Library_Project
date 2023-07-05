package goquery

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

func Weather() {
	url := "https://weather.cma.cn/web/weather/53986.html"

	// 发送HTTP GET请求获取网页内容
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// 使用goquery解析网页内容
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	table := doc.Find("#hourTable_0")

	// 打开或创建一个文件，覆盖已有内容
	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 将元素文本内容写入文件
	_, err = fmt.Fprintln(file, table.Text())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("文本已写入文件")
}
