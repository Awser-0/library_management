package service

import (
	"library/internal/model/do"
	"library/internal/model/entity"
	"library/internal/service/user_service"
	"library/internal/utils/result"
	"time"
)

type IUserService interface {
	UserQuery(query string, page do.QueryPage) do.PageData[entity.User]
	SelectUser(id int64) *entity.User
	SelectUserByUsername(username string) *entity.User
	Login(username, password string) (*entity.User, *result.Result)
	Register(username, password, name, sex, phone string, birth *time.Time, isAdmin bool) (bool, *result.Result)
	UpdateUser(user entity.User) (bool, *result.Result)
}

func NewUserService() IUserService {
	return user_service.New()
}
