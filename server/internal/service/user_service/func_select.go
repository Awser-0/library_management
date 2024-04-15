package user_service

import (
	"library/internal/model/entity"
)

func (s *UserService) SelectUser(id int64) *entity.User {
	return s.userDao.SelectUser(id)
}

func (s *UserService) SelectUserByUsername(username string) *entity.User {
	return s.userDao.SelectUserByUsername(username)
}
