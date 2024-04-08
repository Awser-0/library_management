package userdaoimpl

import (
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (*UserDao) UpdateUser(id int64, user entity.User) bool {
	n, err := g.Model("user").FieldsEx("id", "update_time", "create_time").Data(user).Where("id", id).UpdateAndGetAffected()
	if err != nil {
		panic(err)
	}
	return n == 1
}
