package repository

import (
	m "accounting/src/domain/model"
)

// Department 部門 リポジトリ
type Department interface {
	Save(m.Department) *m.Department
}
