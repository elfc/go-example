package singleton

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSingleton_GetInstance(t *testing.T) {
	srcIns := GetInstance()
	srcIns.Age = 10
	Convey("Given Instance", t, func() {
		ins := GetInstance()

		Convey("When", func() {
			ins.Age = 100

			Convey("Then ShouldResemble", func() {
				So(&srcIns, ShouldResemble, &ins)
			})
		})
	})
}

func BenchmarkSingleton_GetInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		srcIns := GetInstance()
		srcIns.Age = i
	}
}
