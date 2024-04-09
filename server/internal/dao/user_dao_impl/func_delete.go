package userdaoimpl

import (
	"github.com/gogf/gf/v2/frame/g"
)

func (*UserDaoImpl) DeleteUserByID(id int64) bool {
	result, err := g.Model("user").Where("id", id).Delete()
	if err != nil {
		panic(err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return n == 1
}
