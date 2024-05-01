package do

type QueryPage struct {
	PageNum  int `json:"page_num" v:"min:1" d:"1"`
	PageSize int `json:"page_size" v:"min:0" d:"100"`
}
