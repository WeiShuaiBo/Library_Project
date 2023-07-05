package goquery

import "time"

func CalculateDurationToNextTime(now time.Time, hour, minute int) time.Duration {
	// 获取今天08:00的时间点
	nextTime := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())

	// 如果当前时间已经晚于08:00，则计算明天的08:00时间点
	if now.After(nextTime) {
		nextTime = nextTime.Add(24 * time.Hour)
	}

	// 计算距离下一个08:00的时间间隔
	duration := nextTime.Sub(now)

	return duration
}
