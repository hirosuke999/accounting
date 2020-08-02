package model

// AccountTitle 勘定科目
type AccountTitle struct {
	ID   uint
	Code string
	Name string
}

// NewAccountTitle 勘定科目を作成
func NewAccountTitle(code, title string) *AccountTitle {
	return &AccountTitle{
		Code: code,
		Name: title,
	}
}
