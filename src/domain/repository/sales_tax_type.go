package repository

import (
	m "accounting/src/domain/model"
)

// SalesTaxType 消費税区分リポジトリ
type SalesTaxType interface {
	Save(m.SalesTaxType) *m.SalesTaxType
}
