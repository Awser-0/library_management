package user

import "library/internal/service"

type ControllerV1 struct {
	userService service.IUserService
}

func NewV1() ControllerV1 {
	return ControllerV1{
		userService: service.NewUserService(),
	}
}
