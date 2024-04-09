package bookdaoimpl

import "github.com/gogf/gf/v2/frame/g"

func (*BookDaoImpl) DeleteBook(uuid int64) bool {
	result, err := g.Model("book").Where("uuid", uuid).Delete()
	if err != nil {
		panic(err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return n == 1
}
