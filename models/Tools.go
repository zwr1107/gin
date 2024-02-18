package models

import "time"

// TimeToDate 时间戳转时间
func TimeToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// 获取时间戳
func GetTime() int64 {
	return time.Now().Unix()
}

// 获取当前时间
func GetTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
