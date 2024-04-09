package user_service

import (
	"library/internal/model/entity"
	"library/internal/utils/result"
)

func (s *UserService) Login(username, password string) (*entity.User, *result.Result) {
	var user = s.userDao.SelectUserByUsername(username)
	if user == nil {
		return nil, &result.UserNotFound
	}
	if password != user.Password {
		return nil, &result.UserFailPass
	}
	return user, nil
}
