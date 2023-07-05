package app

import (
	"AgroWeather/app/model"
	"AgroWeather/app/router"
	"AgroWeather/app/tools"
)

func Start() {
	defer func() {
		model.Close()
	}()

	model.New()
	tools.NewToken("")

	//// 获取当前时间
	//now := time.Now()
	//// 计算距离下一个08:00的时间间隔
	//duration := goquery.CalculateDurationToNextTime(now, 8, 0)
	//
	//// 创建定时器，在指定的时间间隔后执行方法
	//timer := time.AfterFunc(duration, func() {
	//	goquery.Weather()
	//})
	//
	//<-timer.C

	r := router.New()
	_ = r.Run(":8080")
}
