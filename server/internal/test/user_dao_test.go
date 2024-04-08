package test

import (

	// "library/internal/model/do"
	"fmt"
	"library/internal/dao"
	"library/internal/model/entity"
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/test/gtest"
)

func Test_Select(t *testing.T) {
	var userDao = dao.NewUserDao()
	gtest.C(t, func(t *gtest.T) {
		var user = userDao.SelectUserByUsername("u1")
		fmt.Printf("user: %v\n", user)
		t.Assert(user != nil, true)
	})

	gtest.C(t, func(t *gtest.T) {
		var user = userDao.SelectUserById(1)
		fmt.Printf("user: %v\n", user)
		t.Assert(user != nil, true)
	})
}

func Test_Insert(t *testing.T) {
	var userDao = dao.NewUserDao()
	gtest.C(t, func(t *gtest.T) {
		birth := time.Date(2000, time.Month(2), 2, 15, 30, 20, 0, time.Local)
		t.Assert(userDao.InsertUser(entity.User{
			Username: "u2",
			Password: "1234",
			Sex:      "0",
			Phone:    "12138",
			Birth:    &birth,
			IsAdmin:  false,
		}), true)
	})
}

func Test_Update(t *testing.T) {
	var userDao = dao.NewUserDao()
	gtest.C(t, func(t *gtest.T) {
		var user = userDao.SelectUserByUsername("u1")
		t.Assert(user != nil, true)
		fmt.Printf("user: %v\n", user)
		tb := user.Birth.Add(time.Hour * 1)
		user.Birth = &tb
		fmt.Printf("user: %v\n", user)
		t.Assert(userDao.UpdateUser(user.Id, *user), true)
	})
}

func Test_Delete(t *testing.T) {
	var userDao = dao.NewUserDao()
	gtest.C(t, func(t *gtest.T) {
		var username = "delete_uu"
		birth := time.Date(2000, time.Month(2), 2, 15, 30, 20, 0, time.Local)
		userDao.InsertUser(entity.User{
			Username: username,
			Password: "1234",
			Sex:      "0",
			Phone:    "12138",
			Birth:    &birth,
			IsAdmin:  false,
		})
		user := userDao.SelectUserByUsername(username)
		fmt.Printf("user: %v\n", user)
		t.Assert(userDao.DeleteUserByID(user.Id), true)
	})
}
