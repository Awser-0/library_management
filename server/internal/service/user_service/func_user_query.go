package user_service

import (
	"library/internal/model/do"
	"library/internal/model/entity"
)

func (s *UserService) UserQuery(query string, page do.QueryPage) do.PageData[entity.User] {
	return s.userDao.SelectUsersByQuery(query, page)
}
