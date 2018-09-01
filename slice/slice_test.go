package slice

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSlice_RemoveValue(t *testing.T) {
	var newSlice Slice = NewSlice()

	Convey("Given 一个int64类型切片", t, func() {
		testSlice := []int64{1, 2, 3}

		Convey("When 测试中间值", func() {
			testMiddle := newSlice.RemoveValue(testSlice, 2)

			Convey("Then 不能为空", func() {
				So(testMiddle, ShouldNotBeEmpty)
			})

			Convey("Then ShouldResemble", func() {
				So(testMiddle, ShouldResemble, []int64{1, 3})
			})
		})

		Convey("When 测试头边界值", func() {
			testHead := newSlice.RemoveValue(testSlice, 1)

			Convey("Then 不能为空", func() {
				So(testHead, ShouldNotBeEmpty)
			})

			Convey("Then ShouldResemble", func() {
				So(testHead, ShouldResemble, []int64{2, 3})
			})
		})

		Convey("When 测试尾部边界值", func() {
			testTail := newSlice.RemoveValue(testSlice, 3)

			Convey("Then 不能为空", func() {
				So(testTail, ShouldNotBeEmpty)
			})

			Convey("Then ShouldResemble", func() {
				So(testTail, ShouldResemble, []int64{1, 2})
			})
		})

		Convey("When 测试越界值", func() {
			testUndefined := newSlice.RemoveValue(testSlice, 4)

			Convey("Then 不能为空", func() {
				So(testUndefined, ShouldNotBeEmpty)
			})

			Convey("Then ShouldResemble", func() {
				So(testUndefined, ShouldResemble, []int64{1, 2, 3})
			})
		})
	})
}

func genSlice() (testSlice []int64) {
	N := 1000000
	for i := 0; i < N; i++ {
		testSlice = append(testSlice, int64(i))
	}
	return
}

func BenchmarkSlice_RemoveValue(b *testing.B) {
	s := genSlice()
	var newSlice Slice = NewSlice()
	for i := 0; i < b.N; i++ {
		newSlice.RemoveValue(s, 9999999)
	}
}
