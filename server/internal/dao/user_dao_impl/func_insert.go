package userdaoimpl

import (
	"library/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (*UserDaoImpl) InsertUser(user entity.User) bool {
	result, err := g.Model("user").FieldsEx("id", "update_time", "create_time").Data(user).Insert()
	if err != nil {
		panic(err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return n == 1
}
