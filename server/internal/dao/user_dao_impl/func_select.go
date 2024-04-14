package userdaoimpl

import (
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 通过id选择用户
func (*UserDaoImpl) SelectUser(id int64) *entity.User {
	var user *entity.User
	if result, err := g.Model("user").Where(g.Map{"id": id}).One(); err != nil {
		panic(err)
	} else if err := result.Struct(&user); err != nil {
		panic(err)
	}
	return user
}

// 通过查询选择用户
func (*UserDaoImpl) SelectUsersByQuery(query string) []entity.User {
	query = query + "%"
	var users = []entity.User{}
	if result, err := g.Model("user").WhereLike("username", query).WhereOrLike("profile_name", query).All(); err != nil {
		panic(err)
	} else {
		if err := result.Structs(&users); err != nil {
			panic(err)
		}
	}
	return users
}

// 通过username选择用户
func (*UserDaoImpl) SelectUserByUsername(username string) *entity.User {
	var user *entity.User
	if result, err := g.Model("user").Where(g.Map{"username": username}).One(); err != nil {
		panic(err)
	} else if err := result.Struct(&user); err != nil {
		panic(err)
	}
	return user
}
