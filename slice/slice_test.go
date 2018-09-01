package slice

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSlice_RemoveValue(t *testing.T) {
	var newSlice Slice = NewSlice()

	Convey("初始化一个int64类型切片", t, func() {
		testSlice := []int64{1, 2, 3}

		Convey("测试中间值", func() {
			testMiddle := newSlice.RemoveValue(testSlice, 2)

			Convey("不能为空", func() {
				So(testMiddle, ShouldNotBeEmpty)
			})

			Convey("ShouldResemble", func() {
				So(testMiddle, ShouldResemble, []int64{1, 3})
			})
		})

		Convey("测试头边界值", func() {
			testHead := newSlice.RemoveValue(testSlice, 1)

			Convey("不能为空", func() {
				So(testHead, ShouldNotBeEmpty)
			})

			Convey("ShouldResemble", func() {
				So(testHead, ShouldResemble, []int64{2, 3})
			})
		})

		Convey("测试尾部边界值", func() {
			testTail := newSlice.RemoveValue(testSlice, 3)

			Convey("不能为空", func() {
				So(testTail, ShouldNotBeEmpty)
			})

			Convey("ShouldResemble", func() {
				So(testTail, ShouldResemble, []int64{1, 2})
			})
		})

		Convey("测试越界值", func() {
			testUndefined := newSlice.RemoveValue(testSlice, 4)

			Convey("不能为空", func() {
				So(testUndefined, ShouldNotBeEmpty)
			})

			Convey("ShouldResemble", func() {
				So(testUndefined, ShouldResemble, []int64{1, 2, 3})
			})
		})
	})
}
