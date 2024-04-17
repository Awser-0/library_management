package user_service

import (
	"fmt"
	"library/internal/model/entity"
	"library/internal/utils/result"
)

func (s *UserService) UpdateUser(user entity.User) (bool, *result.Result) {
	var u = s.userDao.SelectUserByUsername(user.Username)
	fmt.Printf("u: %v\n", u)
	fmt.Printf("user: %v\n", user)
	if u != nil && u.Id != user.Id {
		return false, &result.UserExists
	}
	return s.userDao.UpdateUser(user.Id, user), nil
}
