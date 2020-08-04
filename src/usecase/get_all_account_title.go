package usecase

import (
	r "accounting/src/domain/repository"
)

// GetAccountTitleOutput アウトプットパラメータ
type GetAccountTitleOutput struct {
	ID   uint
	Code string
	Name string
}

// GetAllAccountTitle 勘定科目の一覧を取得する
func GetAllAccountTitle(repo r.AccountTitle) []*GetAccountTitleOutput {
	accountTitles := repo.GetAll()

	results := []*GetAccountTitleOutput{}
	for i := range accountTitles {
		results = append(results, &GetAccountTitleOutput{
			ID:   accountTitles[i].ID,
			Code: accountTitles[i].Code,
			Name: accountTitles[i].Name,
		})
	}

	return results
}
