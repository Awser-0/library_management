package service

import (
	"library/internal/model/entity"
	"library/internal/service/user_service"
	"library/internal/utils/result"
	"time"
)

type IUserService interface {
	SelectUser(id int64) *entity.User
	Login(username, password string) (*entity.User, *result.Result)
	Register(username, password, sex, phone string, birth *time.Time, isAdmin bool) (bool, *result.Result)
}

func NewUserService() IUserService {
	return user_service.New()
}
