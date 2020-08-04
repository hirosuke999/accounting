package repository

import (
	m "accounting/src/domain/model"
)

// AccountTitle 勘定科目リポジトリ
type AccountTitle interface {
	Save(m.AccountTitle) *m.AccountTitle
	GetAll() []m.AccountTitle
	GetSalesTaxTypes(ids []m.ID) m.SalesTaxTypes
}
