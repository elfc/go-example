package container

import (
	"container/heap"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &IntHeap{1, 5, 2}

	Convey("初始化堆", t, func() {
		heap.Init(h)

		Convey("When Push", func() {
			heap.Push(h, 8)

			Convey("Then 堆应为", func() {
				hPush := &IntHeap{1, 5, 2, 8}
				So(h.Len(), ShouldResemble, 4)
				So(h, ShouldResemble, hPush)
			})
		})

		Convey("When Pop", func() {
			heap.Pop(h)

			Convey("Then 堆应为", func() {
				hPop := &IntHeap{8, 5, 2}
				So(h.Len(), ShouldResemble, 3)
				So(h, ShouldResemble, hPop)
			})
		})
	})
}
