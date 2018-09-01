package main

import (
	"sort"
)

type Book struct {
	Id   int64
	Name string
}

type bookWrapper struct {
	book []*Book
	by   func(p, q *Book) bool
}

type SortBy func(p, q *Book) bool

func (t bookWrapper) Len() int { // 重写 Len() 方法
	return len(t.book)
}
func (t bookWrapper) Swap(i, j int) { // 重写 Swap() 方法
	t.book[i], t.book[j] = t.book[j], t.book[i]
}
func (t bookWrapper) Less(i, j int) bool { // 重写 Less() 方法
	return t.by(t.book[i], t.book[j])
}

// 封装成 Sortbook 方法
func sortBook(book []*Book, by SortBy) {
	sort.Sort(bookWrapper{book, by})
}
