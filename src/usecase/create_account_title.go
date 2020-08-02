package usecase

import (
	m "accounting/src/domain/model"
	r "accounting/src/domain/repository"
)

// CreateAccountTitleInput インプットパラメータ
type CreateAccountTitleInput struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// CreateAccountTitleOutput アウトプットパラメータ
type CreateAccountTitleOutput struct {
	ID   uint
	Code string
	Name string
}

// CreateAccountTitle 勘定科目を作成する
func CreateAccountTitle(repos r.AccountTitleRepository, input *CreateAccountTitleInput) *CreateAccountTitleOutput {
	accountTitle := m.NewAccountTitle(input.Code, input.Name)
	accountTitle = repos.Save(accountTitle)

	return &CreateAccountTitleOutput{
		ID:   accountTitle.ID,
		Code: accountTitle.Code,
		Name: accountTitle.Name,
	}
}
