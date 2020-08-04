package store

import (
	m "accounting/src/domain/model"

	"github.com/jinzhu/gorm"
)

// AccountTitle 勘定科目ORM
type AccountTitle struct {
	gorm.Model
	Code     string
	Name     string
	TaxTypes SalesTaxTypes `gorm:"many2many:account_titles_sales_tax_types;"`
}

func (orm AccountTitle) toModel() *m.AccountTitle {
	return &m.AccountTitle{ID: orm.ID, Code: orm.Code, Name: orm.Name, TaxTypes: orm.TaxTypes.toModel()}
}

// AccountTitleStore 勘定科目ストア
type AccountTitleStore struct {
	*gorm.DB
}

// NewAccountTitleStore 勘定科目ストアを作成
func NewAccountTitleStore() *AccountTitleStore {
	return &AccountTitleStore{DB: GetDB()}
}

// Save 勘定科目を保存
func (store AccountTitleStore) Save(accountTitle m.AccountTitle) *m.AccountTitle {
	taxTypeIds := accountTitle.GetTaxTypeIds()
	salesTaxTypesOrm := store.getSalesTaxTypes(taxTypeIds)
	orm := &AccountTitle{Code: accountTitle.Code, Name: accountTitle.Name}

	store.DB.Create(orm)
	store.DB.Model(orm).Association("TaxTypes").Append(salesTaxTypesOrm)

	return orm.toModel()
}

// GetAll 勘定科目の一覧を取得
func (store AccountTitleStore) GetAll() (accountTitles []m.AccountTitle) {
	var accountTitlesOrm []AccountTitle
	store.DB.Find(&accountTitlesOrm)

	for i := range accountTitlesOrm {
		accountTitles = append(accountTitles, m.AccountTitle{
			ID:   accountTitlesOrm[i].ID,
			Code: accountTitlesOrm[i].Code,
			Name: accountTitlesOrm[i].Name,
		})
	}

	return
}

// GetSalesTaxTypes 消費税区分の一覧を取得
func (store AccountTitleStore) GetSalesTaxTypes(ids []m.ID) (taxTypes m.SalesTaxTypes) {
	salesTaxTypesOrm := store.getSalesTaxTypes(ids)

	for i := range salesTaxTypesOrm {
		taxTypes = append(taxTypes, m.SalesTaxType{
			ID:   salesTaxTypesOrm[i].ID,
			Name: salesTaxTypesOrm[i].Name,
		})
	}

	return
}

func (store AccountTitleStore) getSalesTaxTypes(ids []m.ID) (taxTypes []SalesTaxType) {
	store.DB.Where(ids).Find(&taxTypes)
	return
}
