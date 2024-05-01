package do

type PageData[T any] struct {
	List     []T `json:"list"`
	PageNum  int `json:"page_num"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}
