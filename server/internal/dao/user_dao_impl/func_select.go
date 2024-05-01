package userdaoimpl

import (
	"library/internal/model/do"
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
func (*UserDaoImpl) SelectUsersByQuery(query string, page do.QueryPage) do.PageData[entity.User] {
	query = query + "%"
	var users = make([]entity.User, 0)
	var total int
	if err := g.Model("user").WhereLike("username", query).WhereOrLike("nickname", query).Limit((page.PageNum-1)*(page.PageSize), page.PageSize).WithAll().ScanAndCount(&users, &total, false); err != nil {
		panic(err)
	}
	return do.PageData[entity.User]{
		List:     users,
		PageNum:  page.PageNum,
		PageSize: page.PageSize,
		Total:    total,
	}
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
