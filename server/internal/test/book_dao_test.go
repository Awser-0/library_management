package test

import (
	"fmt"
	"library/internal/dao"
	"library/internal/model/entity"
	"testing"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_BookDao_Insert(t *testing.T) {
	var dao = dao.NewBookDao()
	gtest.C(t, func(t *gtest.T) {
		t.Assert(dao.InsertBook(entity.Book{
			Title:        "白夜行",
			Author:       "东野圭吾",
			Cover:        "cover.png",
			Introduction: "简介",
			PublishTime:  nil,
			Total:        20,
		}), true)
	})
}

func Test_BookDao_Select(t *testing.T) {
	var dao = dao.NewBookDao()
	// gtest.C(t, func(t *gtest.T) {
	// 	var book = dao.SelectBook(10000001)
	// 	fmt.Printf("book: %v\n", book)
	// 	t.Assert(book != nil, true)
	// })
	gtest.C(t, func(t *gtest.T) {
		books := dao.SelectBooksRestricted("东野圭吾")
		for i, book := range books {
			fmt.Printf("%v: %v\n", i, book)
		}
	})
}

func Test_BookDao_Update(t *testing.T) {
	var dao = dao.NewBookDao()
	gtest.C(t, func(t *gtest.T) {
		var book = dao.SelectBook(10000001)
		book.Introduction = "修改简介1"
		t.Assert(dao.UpdateBook(book.UUID, *book), true)
	})
}

func Test_BookDao_Delete(t *testing.T) {
	var dao = dao.NewBookDao()
	gtest.C(t, func(t *gtest.T) {
		t.Assert(dao.DeleteBook(10000004), true)
	})
}
