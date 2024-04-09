package dao

import (
	userdaoimpl "library/internal/dao/user_dao_impl"
	"library/internal/model/entity"
)

type IUserDao interface {
	SelectUserById(id int64) *entity.User
	SelectUserByUsername(username string) *entity.User
	InsertUser(user entity.User) bool
	UpdateUser(id int64, user entity.User) bool
	DeleteUserByID(id int64) bool
}

func NewUserDao() IUserDao {
	return userdaoimpl.New()
}
