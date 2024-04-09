package do

type BorrowRecordQueryCondition struct {
	UserId   int64  `json:"user_id"`
	BookUUID int64  `json:"book_uuid"`
	State    string `json:"state"`
}
