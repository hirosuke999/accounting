package model

import "accounting/src/domain/model/enum"

// AccountTitle 勘定科目
type AccountTitle struct {
	ID                       ID
	Code                     string
	Name                     string
	TaxTypes                 SalesTaxTypes
	DepartmentNecessity      enum.DepartmentNecessity
	SegmentNecessity         enum.SegmentNecessity
	ProjectNecessity         enum.ProjectNecessity
	BusinessPartnerNecessity enum.BusinessPartnerNecessity
	ReasonNoteNecessity      enum.ReasonNoteNecessity
}

// NewAccountTitle 勘定科目を作成
func NewAccountTitle(code, title string, taxTypes SalesTaxTypes) *AccountTitle {
	return &AccountTitle{
		Code:                     code,
		Name:                     title,
		TaxTypes:                 taxTypes,
		DepartmentNecessity:      enum.DepartmentNecessityNone,
		SegmentNecessity:         enum.SegmentNecessityNone,
		ProjectNecessity:         enum.ProjectNecessityNone,
		BusinessPartnerNecessity: enum.BusinessPartnerNecessityNone,
		ReasonNoteNecessity:      enum.ReasonNoteNecessityNone,
	}
}

// GetTaxTypeIds 消費税区分コレクションのID一覧を取得
func (title AccountTitle) GetTaxTypeIds() []ID {
	return title.TaxTypes.GetIds()
}
