package timelib

import (
	"strconv"
	"time"
)

// TimeNowFormat 日期格式
func TimeNowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// UnixTimeFormat 解析unixtime
func UnixTimeFormat(timestamp int64) string {
	utime := time.Unix(timestamp, 0)
	return utime.Format("2006-01-02 15:04:05")
}

// StrTimeFormat unix 字符串转换日期
func StrTimeFormat(timestamp string) string {
	ui, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	utime := time.Unix(ui, 0)
	return utime.Format("2006-01-02 15:04:05")
}
