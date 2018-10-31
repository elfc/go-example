package util

import (
	"errors"
	"time"
)

const (
	ERROR_TIMEUTIL_LOCATION = "The two time location are different"
	ERROR_TIMEUTIL_GT       = "The time1 Should be gt than time2"
)

const (
	TempFormat    = "2006-01-02 15:04:05"
	TempDayFormat = "2006-01-02"
)

type TimeUtil interface {
	// 字符串转时间
	ParseStrTime(str string) time.Time
	// unix 时间戳转换为时间格式
	ParseUnixTime(sr int64) time.Time
	// 转换天格式的字符串为时间格式
	ParseDayStrTime(str string) time.Time
	// 字符串时间格式, 转换为时间格式并且增加8小时
	StrTimeAdd8H(str string) string
	// 添加时间
	ParseAddDateToStr(s time.Time, years, months, days int) string
	// 把时间转换为字符串
	ParseTimeToStr(s time.Time) string
	// 计算t1与t2之间日时间差, t1 一定要大于t2
	// 小于24小时时间差为0天
	SubDay(t1, t2 time.Time) (int, error)
}

func NewTime() TimeUtil {
	return new(timeUtil)
}

func newTimeUtil() TimeUtil {
	return new(timeUtil)
}

// 用于默认引入调用
var NewTimeUtil TimeUtil = newTimeUtil()

type timeUtil struct{}

func (t *timeUtil) ParseStrTime(str string) time.Time {
	parseStrTime, _ := time.ParseInLocation(TempFormat, str, time.Local)
	return parseStrTime
}

func (t *timeUtil) ParseDayStrTime(str string) time.Time {
	parseStrTime, _ := time.ParseInLocation(TempDayFormat, str, time.Local)
	return parseStrTime
}

func (t *timeUtil) ParseUnixTime(sr int64) time.Time {
	return time.Unix(sr, 0).Local()
}

func (t *timeUtil) StrTimeAdd8H(str string) string {
	parseStrTime := t.ParseStrTime(str)
	h, _ := time.ParseDuration("8h")
	afterTime := parseStrTime.Add(h)

	return afterTime.Format(TempFormat)
}

func (t *timeUtil) ParseAddDateToStr(s time.Time, years, months, days int) string {
	adate := s.AddDate(years, months, days)
	return adate.Format(TempFormat)
}

func (t *timeUtil) ParseTimeToStr(s time.Time) string {
	return s.Format(TempFormat)
}

func (t *timeUtil) SubDay(t1, t2 time.Time) (int, error) {
	var day int
	if t1.Unix() < t2.Unix() {
		return 0, errors.New(ERROR_TIMEUTIL_GT)
	}

	if t1.Location().String() != t2.Location().String() {
		return 0, errors.New(ERROR_TIMEUTIL_LOCATION)
	}

	hours := t1.Sub(t2).Hours()

	tempDay := hours / 24

	if tempDay < 1 {
		day = 0
	}

	day = int(tempDay)

	return day, nil
}
