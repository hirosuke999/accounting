package store

import (
	m "accounting/src/domain/model"

	"github.com/jinzhu/gorm"
)

// AccountTitleORM 勘定科目ORM
type AccountTitleORM struct {
	gorm.Model
	Code string
	Name string
}

// AccountTitleStore 勘定科目ストア
type AccountTitleStore struct {
	DB *gorm.DB
}

// NewAccountTitleStore 勘定科目ストアを作成
func NewAccountTitleStore(db *gorm.DB) *AccountTitleStore {
	return &AccountTitleStore{DB: db}
}

// Save 勘定科目を保存
func (tbl AccountTitleStore) Save(accountTitle *m.AccountTitle) *m.AccountTitle {
	orm := &AccountTitleORM{Code: accountTitle.Code, Name: accountTitle.Name}
	tbl.DB.Create(orm)
	return &m.AccountTitle{ID: orm.ID, Code: orm.Code, Name: orm.Name}
}

// GetAll 勘定科目の一覧を取得
func (tbl AccountTitleStore) GetAll() []*m.AccountTitle {
	accountTitlesOrm := []AccountTitleORM{}
	tbl.DB.Find(&accountTitlesOrm)

	accountTitles := []*m.AccountTitle{}
	for i := range accountTitlesOrm {
		accountTitles = append(accountTitles, &m.AccountTitle{
			ID:   accountTitlesOrm[i].ID,
			Code: accountTitlesOrm[i].Code,
			Name: accountTitlesOrm[i].Name,
		})
	}

	return accountTitles
}
