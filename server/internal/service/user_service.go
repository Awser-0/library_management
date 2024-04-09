package service

import "library/internal/service/user_service"

type IUserService interface{}

func NewUserService() IUserService {
	return user_service.New()
}
