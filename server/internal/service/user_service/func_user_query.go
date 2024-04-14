package user_service

import "library/internal/model/entity"

func (s *UserService) UserQuery(query string) []entity.User {
	return s.userDao.SelectUsersByQuery(query)
}
