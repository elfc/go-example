package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
	"time"
)

func TestTimeUtil_ParseStrTime(t *testing.T) {
	Convey("Given 一个时间字符串", t, func() {

		Convey("When 测试字符串转换为时间", func() {
			testTimeStr := "2018-10-09 14:25:05"

			parseStrTime := NewTimeUtil.ParseStrTime(testTimeStr)

			Convey("Then ShouldEqual", func() {
				So(reflect.TypeOf(parseStrTime), ShouldEqual, reflect.TypeOf(time.Now()))
			})
		})

		Convey("When 测试字符串格式为错误时", func() {
			testTimeStr := "12345"

			parseStrTime := NewTimeUtil.ParseStrTime(testTimeStr)

			Convey("Then ShouldEqual", func() {
				So(reflect.TypeOf(parseStrTime), ShouldEqual, reflect.TypeOf(time.Now()))
			})
		})
	})
}

func TestTimeUtil_ParseDayStrTime(t *testing.T) {
	Convey("Given 一个时间字符串", t, func() {

		Convey("When 测试字符串转换为时间", func() {
			testTimeStr := "2018-10-09"

			parseStrTime := NewTimeUtil.ParseDayStrTime(testTimeStr)

			Convey("Then ShouldEqual", func() {
				So(reflect.TypeOf(parseStrTime), ShouldEqual, reflect.TypeOf(time.Now()))
			})
		})
	})
}

func TestTimeUtil_ParseUnixTime(t *testing.T) {
	Convey("Given 一个unix时间戳", t, func() {
		testTimeUnix := int64(1538984415000)

		Convey("When 测试时间戳转换为时间", func() {
			parseUnixTime := NewTimeUtil.ParseUnixTime(testTimeUnix)

			Convey("Then ShouldEqual", func() {
				So(reflect.TypeOf(parseUnixTime), ShouldEqual, reflect.TypeOf(time.Now()))
			})
		})
	})
}

func TestTimeUtil_StrTimeAdd8H(t *testing.T) {
	Convey("Given 一个时间字符串", t, func() {
		testTimeStr := "2018-10-09 14:25:05"

		Convey("When 测试字符串转换为时间并且增加8小时再转回字符串", func() {
			parseStrTime := NewTimeUtil.StrTimeAdd8H(testTimeStr)

			Convey("Then ShouldEqual", func() {
				So(parseStrTime, ShouldEqual, "2018-10-09 22:25:05")
			})
		})
	})
}

func TestTimeUtil_ParseAddDateToStr(t *testing.T) {
	Convey("Given 一个时间字符串", t, func() {
		testTimeStr := "2018-10-09 00:00:00"

		Convey("When 测试时间增加一天", func() {
			testTime, _ := time.Parse(TempFormat, testTimeStr)

			addDateTime := NewTimeUtil.ParseAddDateToStr(testTime, 0, 0, 1)
			Convey("Then ShouldEqual", func() {
				So(addDateTime, ShouldEqual, "2018-10-10 00:00:00")
			})
		})
	})
}

func TestTimeUtil_ParseTimeToStr(t *testing.T) {
	Convey("Given 一个时间", t, func() {
		testTime := time.Now()

		Convey("When 测试时间转换为字符串", func() {
			parseTimeToStr := NewTimeUtil.ParseTimeToStr(testTime)

			Convey("Then ShouldEqual", func() {
				So(reflect.TypeOf(parseTimeToStr), ShouldEqual, reflect.TypeOf("2018-10-10 00:00:00"))
			})
		})
	})
}

func TestTimeUtil_SubDay(t *testing.T) {
	Convey("Given 测试时间差", t, func() {
		Convey("When 测试一天时间差", func() {
			testTimeStr1 := "2018-10-11 00:00:00"
			testTimeStr2 := "2018-10-10 00:00:00"
			testTime1, _ := time.ParseInLocation(TempFormat, testTimeStr1, time.Local)
			testTime2, _ := time.ParseInLocation(TempFormat, testTimeStr2, time.Local)
			subTime, err := NewTimeUtil.SubDay(testTime1, testTime2)

			Convey("Then ShouldBeNil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then ShouldResemble", func() {
				So(subTime, ShouldResemble, 1)
			})
		})

		Convey("When 测试大于2天小于3天时间差", func() {
			testTimeStr1 := "2018-10-12 03:00:00"
			testTimeStr2 := "2018-10-10 00:00:00"
			testTime1, _ := time.ParseInLocation(TempFormat, testTimeStr1, time.Local)
			testTime2, _ := time.ParseInLocation(TempFormat, testTimeStr2, time.Local)
			subTime, err := NewTimeUtil.SubDay(testTime1, testTime2)

			Convey("Then ShouldBeNil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then ShouldResemble", func() {
				So(subTime, ShouldResemble, 2)
			})
		})

		Convey("When 测试小于一天时间差", func() {
			testTimeStr1 := "2018-10-10 15:00:00"
			testTimeStr2 := "2018-10-10 00:00:00"
			testTime1, _ := time.ParseInLocation(TempFormat, testTimeStr1, time.Local)
			testTime2, _ := time.ParseInLocation(TempFormat, testTimeStr2, time.Local)
			subTime, err := NewTimeUtil.SubDay(testTime1, testTime2)

			Convey("Then ShouldBeNil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then ShouldResemble", func() {
				So(subTime, ShouldResemble, 0)
			})
		})

		Convey("When 测试location 不同", func() {
			testTimeStr1 := "2018-10-10 15:00:00"
			testTimeStr2 := "2018-10-10 00:00:00"
			testTime1, _ := time.ParseInLocation(TempFormat, testTimeStr1, time.Local)
			testTime2, _ := time.ParseInLocation(TempFormat, testTimeStr2, time.Local)
			_, err := NewTimeUtil.SubDay(testTime1.UTC(), testTime2)

			Convey("Then ShouldNotBeNil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then ShouldResemble", func() {
				So(err.Error(), ShouldResemble, ERROR_TIMEUTIL_LOCATION)
			})
		})

		Convey("When 测试t1 时间小于t2 时间", func() {
			testTimeStr1 := "2018-10-09 15:00:00"
			testTimeStr2 := "2018-10-10 00:00:00"
			testTime1, _ := time.ParseInLocation(TempFormat, testTimeStr1, time.Local)
			testTime2, _ := time.ParseInLocation(TempFormat, testTimeStr2, time.Local)
			_, err := NewTimeUtil.SubDay(testTime1, testTime2)

			Convey("Then ShouldNotBeNil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then ShouldResemble", func() {
				So(err.Error(), ShouldResemble, ERROR_TIMEUTIL_GT)
			})
		})
	})
}
