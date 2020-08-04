package usecase

import (
	m "accounting/src/domain/model"
	"accounting/src/domain/model/enum"
	"accounting/src/domain/repository"
)

// CreateAccountTitleInput インプットデータ構造
type CreateAccountTitleInput struct {
	Code                     string                        `json:"code"`
	Name                     string                        `json:"name"`
	TaxTypeIds               []m.ID                        `json:"tax_type_ids"`
	DepartmentNecessity      enum.DepartmentNecessity      `json:"department_necessity"`
	SegmentNecessity         enum.SegmentNecessity         `json:"segment_necessity"`
	ProjectNecessity         enum.ProjectNecessity         `json:"project_necessity"`
	BusinessPartnerNecessity enum.BusinessPartnerNecessity `json:"business_partner_necessity"`
	ReasonNoteNecessity      enum.ReasonNoteNecessity      `json:"reason_note_necessity"`
}

// CreateAccountTitleOutput アウトプットデータ構造
type CreateAccountTitleOutput struct {
	ID                       m.ID
	Code                     string
	Name                     string
	TaxTypes                 m.SalesTaxTypes
	DepartmentNecessity      enum.DepartmentNecessity
	SegmentNecessity         enum.SegmentNecessity
	ProjectNecessity         enum.ProjectNecessity
	BusinessPartnerNecessity enum.BusinessPartnerNecessity
	ReasonNoteNecessity      enum.ReasonNoteNecessity
}

// CreateAccountTitle 勘定科目を作成する
func CreateAccountTitle(repo repository.AccountTitle, in CreateAccountTitleInput) *CreateAccountTitleOutput {
	salesTaxTypes := repo.GetSalesTaxTypes(in.TaxTypeIds)
	accountTitle := m.NewAccountTitle(in.Code, in.Name, salesTaxTypes)
	accountTitle = repo.Save(*accountTitle)

	return &CreateAccountTitleOutput{
		ID:                       accountTitle.ID,
		Code:                     accountTitle.Code,
		Name:                     accountTitle.Name,
		TaxTypes:                 accountTitle.TaxTypes,
		DepartmentNecessity:      accountTitle.DepartmentNecessity,
		SegmentNecessity:         accountTitle.SegmentNecessity,
		ProjectNecessity:         accountTitle.ProjectNecessity,
		BusinessPartnerNecessity: accountTitle.BusinessPartnerNecessity,
		ReasonNoteNecessity:      accountTitle.ReasonNoteNecessity,
	}
}
