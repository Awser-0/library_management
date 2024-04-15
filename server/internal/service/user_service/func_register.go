package user_service

import (
	"library/internal/model/entity"
	"library/internal/utils/result"
	"time"
)

func (s *UserService) Register(username, password, name, sex, phone string, birth *time.Time, isAdmin bool) (bool, *result.Result) {
	if s.userDao.SelectUserByUsername(username) != nil {
		return false, &result.UserExists
	}
	if sex != entity.UserSexGirl {
		sex = entity.UserSexBoy
	}
	var ok = s.userDao.InsertUser(entity.User{
		Username: username,
		Password: password,
		Name:     name,
		Sex:      sex,
		Phone:    phone,
		Birth:    birth,
		IsAdmin:  isAdmin,
	})
	return ok, nil
}
