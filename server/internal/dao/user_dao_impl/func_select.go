package userdaoimpl

import (
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (*UserDaoImpl) SelectUserById(id int64) *entity.User {
	var user *entity.User
	if result, err := g.Model("user").Where(g.Map{"id": id}).One(); err != nil {
		panic(err)
	} else if err := result.Struct(&user); err != nil {
		panic(err)
	}
	return user
}

func (*UserDaoImpl) SelectUserByUsername(username string) *entity.User {
	var user *entity.User
	if result, err := g.Model("user").Where(g.Map{"username": username}).One(); err != nil {
		panic(err)
	} else if err := result.Struct(&user); err != nil {
		panic(err)
	}
	return user
}
