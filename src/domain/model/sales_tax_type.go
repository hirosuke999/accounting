package model

// SalesTaxType 消費税区分
type SalesTaxType struct {
	ID   ID
	Name string
}

// NewSalesTaxType 消費税区分を作成
func NewSalesTaxType(name string) *SalesTaxType {
	return &SalesTaxType{Name: name}
}

// SalesTaxTypes 消費税区分コレクション
type SalesTaxTypes []SalesTaxType

// GetIds 消費税区分コレクションのID一覧を取得
func (taxTypes SalesTaxTypes) GetIds() (ids []ID) {
	for i := range taxTypes {
		ids = append(ids, taxTypes[i].ID)
	}
	return
}
