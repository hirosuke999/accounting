package repository

import (
	m "accounting/src/domain/model"
)

// AccountTitleRepository 勘定科目リポジトリ
type AccountTitleRepository interface {
	Save(*m.AccountTitle) *m.AccountTitle
	GetAll() []*m.AccountTitle
}
