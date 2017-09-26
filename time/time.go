package time

import (
	"time"
	"thelark.cn/golang.ext/base"
)

// 标准
const (
	yyyy = "2006" // 年
	MM   = "01"   // 月
	dd   = "02"   // 日
	HH   = "15"   // 时
	mm   = "04"   // 分
	ss   = "05"   // 秒
)
const (
	yy  = "06" // 年
	星期  = "Mon"
	月   = "Jan"
	日   = "_2"
	时区  = "MST"
	纳秒  = ".000"
	N纳秒 = ".000000"
)

const (
	ANSIC    = "Mon Jan _2 15:04:05 2006"
	UNIXDATE = "Mon Jan _2 15:04:05 MST 2006"
	RFC822   = "Mon Jan _2 15:04:05 MST 2006"
)
const SHORT_TIME_LAYOUT = "2006-01-02 15:04:05"
const LONG_TIME_LAYOUT = "2006-01-02 15:04:05.000"
const LONGL_TIME_LAYOUT = "2006-01-02 15:04:05.000000"

const DEFAULT_TIME_LAYOUT = "2006-01-02 15:04:05"

/**
 * TODO: time.Time => string
 * @param time 时间
 * @param timeLayout 时间格式字符串
 * @return string
 */
func FormatTime(time time.Time, timeLayout string) string {
	if base.StringIsNull(timeLayout) {
		timeLayout = DEFAULT_TIME_LAYOUT
	}
	return time.Format(timeLayout)
}

/**
 * TODO: 时间戳 => string
 * @param timestamp 时间戳
 * @param timeLayout 时间格式字符串
 * @return string
 */
func FormatTimestamp(timestamp int64, timeLayout string) string {
	if base.StringIsNull(timeLayout) {
		timeLayout = DEFAULT_TIME_LAYOUT
	}
	return time.Unix(timestamp, 0).Format(timeLayout)
}

/**
 * TODO: string => (UTC)time.Time
 * @param s time.Time的字符串形式
 * @param timeLayout 时间格式字符串
 * @return (UTC)time.Time
 */
func ParseTime(s string, timeLayout string) time.Time {
	if base.StringIsNull(timeLayout) {
		timeLayout = DEFAULT_TIME_LAYOUT
	}
	theTime, _ := time.Parse(timeLayout, s) //使用模板在对应时区转化为time.time类型
	return theTime
}

/**
 * TODO: string => (本地)time.Time
 * @param s time.Time的字符串形式
 * @param timeLayout 时间格式字符串
 * @return (本地)time.Time
 */
func ParseLocalTime(s string, timeLayout string) time.Time {
	if base.StringIsNull(timeLayout) {
		timeLayout = DEFAULT_TIME_LAYOUT
	}
	local, _ := time.LoadLocation("Local")                   //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, s, local) //使用模板在对应时区转化为time.time类型
	return theTime
}

/**
 * TODO: string => (UTC)时间戳
 * @param s time.Time的字符串形式
 * @param flag 是否转换为 UnixNano 时间戳
 * @param timeLayout 时间格式字符串
 * @return (UTC)时间戳
 */
func ParseTimestamp(s string, flag bool, timeLayout string) int64 {
	if base.StringIsNull(timeLayout) {
		timeLayout = DEFAULT_TIME_LAYOUT
	}
	if flag {
		return ParseTime(s, timeLayout).UnixNano()
	}
	return ParseTime(s, timeLayout).Unix()
}

/**
 * TODO: string => (本地)时间戳
 * @param s time.Time的字符串形式
 * @param flag 是否转换为 UnixNano 时间戳
 * @param timeLayout 时间格式字符串
 * @return (本地)时间戳
 */
func ParseLocalTimestamp(s string, flag bool, timeLayout string) int64 {
	if base.StringIsNull(timeLayout) {
		timeLayout = DEFAULT_TIME_LAYOUT
	}
	if flag {
		return ParseLocalTime(s, timeLayout).UnixNano()
	}
	return ParseLocalTime(s, timeLayout).Unix()
}
