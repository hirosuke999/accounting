package store

import (
	m "accounting/src/domain/model"

	"github.com/jinzhu/gorm"
)

// SalesTaxType 消費税区分ORM
type SalesTaxType struct {
	gorm.Model
	Name          string
	AccountTitles []*AccountTitle `gorm:"many2many:account_titles_sales_tax_types;"`
}

// SalesTaxTypes 消費税区分ORMコレクション
type SalesTaxTypes []*SalesTaxType

func (taxTypes SalesTaxTypes) toModel() (results m.SalesTaxTypes) {
	for i := range taxTypes {
		results = append(results, m.SalesTaxType{
			ID:   taxTypes[i].ID,
			Name: taxTypes[i].Name,
		})
	}
	return
}

// SalesTaxTypeStore 消費税区分ストア
type SalesTaxTypeStore struct {
	*gorm.DB
}

// NewSalesTaxTypeStore 消費税区分ストアを作成
func NewSalesTaxTypeStore() *SalesTaxTypeStore {
	return &SalesTaxTypeStore{DB: GetDB()}
}

// Save 消費税区分を保存
func (store SalesTaxTypeStore) Save(salesTaxType m.SalesTaxType) *m.SalesTaxType {
	orm := &SalesTaxType{Name: salesTaxType.Name}
	store.DB.Create(orm)
	return &m.SalesTaxType{ID: orm.ID, Name: orm.Name}
}
