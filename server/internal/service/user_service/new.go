package user_service

import "library/internal/dao"

type UserService struct {
	userDao dao.IUserDao
}

func New() *UserService {
	return &UserService{
		userDao: dao.NewUserDao(),
	}
}
