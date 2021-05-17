package utils

import "time"

//获取当前年月
func DateNowFormatYm() string {
	tm := time.Now()
	return tm.Format("2006-01")
}

// 获取年月日
func DateNowFormatYmd() string {
	tm := time.Now()
	return tm.Format("2006-01-02")
}

//获取年月日时分（字符串类型）
func DateNowFormatYmdhm() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04")
}

//获取年月日时分秒（字符串类型）
func DateNowFormatYmdhms() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}

//时间戳转年月日
func DateFormatYm(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01")
}

//时间戳转年月日
func DateFormatYmd(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02")
}

// 时间戳转年月日 时分
func DateFormatYmdhm(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02 15:04")
}

// 时间戳转年月日 时分秒
func DateFormatYmdhms(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02 15:04:02")
}

//时间戳
func DateUnix() int {
	t := time.Now().Unix()
	return int(t)
}
//时间戳
func DateUnixNano() int {
	t := time.Now().UnixNano()
	return int(t)
}
// 获取日期的年月日
func DateYmd() (int, int, int) {
	timeNow := time.Now()
	year, month, day := timeNow.Date()
	return year, int(month), day
}

//获取第几周
func DateWeek() int {
	_, week := time.Now().ISOWeek()
	return week
}







