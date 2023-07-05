package goquery

import (
	"AgroWeather/app/model"
	"bufio"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
	"os"
	"regexp"
	"strings"
)

func Preprocess() *model.Weather {

	// 打开文件
	file, err := os.Open("./app/goquery/output.txt")

	// 打开文件
	//file, err := os.Open("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 使用bufio.Scanner读取文件内容
	scanner := bufio.NewScanner(file)

	//预处理

	// 自定义过滤函数，去除空格和汉字
	filter := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 使用正则表达式去除汉字和空格
		re := regexp.MustCompile(`[\p{Han}\s]+`)
		index := re.FindIndex(data)

		if index != nil {
			// 去除汉字和空格
			trimmedData := strings.TrimSpace(string(data[:index[0]]))
			return index[1], []byte(trimmedData), nil
		}

		if !atEOF {
			return 0, nil, nil
		}

		if len(data) == 0 {
			return 0, nil, nil
		}

		return len(data), data, bufio.ErrFinalToken
	}

	scanner.Split(filter)

	//创建一个切片
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Println(lines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 打开或创建一个文件，覆盖已有内容
	file, err2 := os.OpenFile("./app/goquery/output2.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	//file, err2 := os.OpenFile("output2.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err2 != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 将元素文本内容写入文件
	_, err = fmt.Fprintln(file, lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("文本已写入文件")

	time := []interface{}{lines[1], lines[2], lines[3], lines[4], lines[5], lines[6], lines[7], lines[8]}
	temp := []interface{}{lines[9], lines[10], lines[11], lines[12], lines[13], lines[14], lines[15], lines[16]}
	humidity := []interface{}{lines[33], lines[34], lines[35], lines[36], lines[37], lines[38], lines[39], lines[40]}
	cloud := []interface{}{lines[41], lines[42], lines[43], lines[44], lines[45], lines[46], lines[47], lines[47], lines[48]}

	timeStr := make([]string, len(time))
	for i, v := range time {
		timeStr[i] = v.(string)
	}

	tempStr := make([]string, len(temp))
	for i, v := range temp {
		tempStr[i] = v.(string)
	}

	humidityStr := make([]string, len(humidity))
	for i, v := range humidity {
		humidityStr[i] = v.(string)
	}

	cloudStr := make([]string, len(cloud))
	for i, v := range cloud {
		cloudStr[i] = v.(string)
	}

	ret := &model.Weather{}
	ret.Time = timeStr
	ret.Temp = tempStr
	ret.Humidity = humidityStr
	ret.Cloud = cloudStr

	fmt.Println("时间：", time)
	fmt.Println("气温：", temp)
	fmt.Println("湿度：", humidity)
	fmt.Println("云量：", cloud)
	return ret

	//fmt.Println("时间：", time)
	//fmt.Println("气温：", temp)
	//fmt.Println("湿度：", humidity)
	//fmt.Println("云量：", cloud)
	//
	//// 创建一个新的折线图
	//line := charts.NewLine()
	//
	//// 设置折线图的标题和Y轴名称
	//line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Temperature Line Chart"}))
	//line.SetGlobalOptions(charts.WithYAxisOpts(opts.YAxis{Name: "Temperature (°C)"}))
	//
	//// 添加数据到折线图
	//line.SetXAxis(time)
	//line.AddSeries("Temperature", generateLineItems(timeStr, tempFloats))

}

// 辅助函数用于生成折线图数据项
func generateLineItems(months []string, temperatures []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i, temp := range temperatures {
		items = append(items, opts.LineData{Value: temp, Name: months[i]})
	}
	return items
}
