package usecase

import (
	"accounting/src/domain/model"
	m "accounting/src/domain/model"
	r "accounting/src/domain/repository"
)

// CreateSalesTaxTypeInput インプットデータ構造
type CreateSalesTaxTypeInput struct {
	Name string
}

// CreateSalesTaxTypeOutput アウトプットデータ構造
type CreateSalesTaxTypeOutput struct {
	ID   m.ID
	Name string
}

// CreateSalesTaxType 消費税区分を作成する
func CreateSalesTaxType(repo r.SalesTaxType, in CreateSalesTaxTypeInput) *CreateSalesTaxTypeOutput {
	salesTaxType := model.NewSalesTaxType(in.Name)
	salesTaxType = repo.Save(*salesTaxType)

	return &CreateSalesTaxTypeOutput{ID: salesTaxType.ID, Name: salesTaxType.Name}
}
