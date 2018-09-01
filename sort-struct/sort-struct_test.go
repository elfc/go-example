package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

func TestSortStruct(t *testing.T) {
	Convey("Given []*Book", t, func() {
		books := []*Book{
			{
				Id: 3,
			}, {
				Id: 1,
			}, {
				Id: 2,
			},
		}

		Convey("When 正序", func() {
			sort.Sort(bookWrapper{books, func(p, q *Book) bool {
				return p.Id < q.Id
			}})
			expect := []*Book{
				{
					Id: 1,
				}, {
					Id: 2,
				}, {
					Id: 3,
				},
			}
			Convey("Then ShouldResemble", func() {
				So(books, ShouldResemble, expect)
			})
		})

		Convey("When 倒序", func() {
			sort.Sort(bookWrapper{books, func(p, q *Book) bool {
				return p.Id > q.Id
			}})
			expect := []*Book{
				{
					Id: 3,
				}, {
					Id: 2,
				}, {
					Id: 1,
				},
			}
			Convey("Then ShouldResemble", func() {
				So(books, ShouldResemble, expect)
			})
		})
	})
}
