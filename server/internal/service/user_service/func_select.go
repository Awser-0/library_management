package user_service

import (
	"library/internal/model/entity"
)

func (s *UserService) SelectUser(id int64) *entity.User {
	return s.userDao.SelectUserById(id)
}
