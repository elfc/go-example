package iterator

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIterator(t *testing.T) {
	ints := Ints{1, 2, 3}

	Convey("迭代器", t, func() {
		index := 1
		for v := range ints.Iteracor() {
			So(v, ShouldResemble, index)
			index++
		}
	})
}
