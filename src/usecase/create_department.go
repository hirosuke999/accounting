package usecase

import (
	"accounting/src/domain/model"
	"accounting/src/domain/model/value"
	"accounting/src/domain/repository"
)

// CreateDepartmentInput 入力パラメータ
type CreateDepartmentInput struct {
	Code            string       `json:"code"`
	Name            string       `json:"name"`
	IsPersistent    bool         `json:"is_parsistent"`
	ValidPeriod     value.Period `json:"valid_period"`
	IsRecordEnabled bool         `json:"is_record_enabled"`
}

// CreateDepartmentOutput 出力パラメータ
type CreateDepartmentOutput struct {
	ID              model.ID
	Code            string
	Name            string
	IsPersistent    bool
	ValidPeriod     value.Period
	IsRecordEnabled bool
}

// CreateDepartment 部門を作成する
func CreateDepartment(repo repository.Department, in *CreateDepartmentInput) *CreateDepartmentOutput {
	dep := model.NewDepartment(&model.DepartmentParam{
		Code:            in.Code,
		Name:            in.Name,
		IsPersistent:    in.IsPersistent,
		ValidPeriod:     in.ValidPeriod,
		IsRecordEnabled: in.IsRecordEnabled,
	})
	dep = repo.Save(*dep)
	return &CreateDepartmentOutput{
		ID:              dep.ID(),
		Code:            dep.Code(),
		Name:            dep.Name(),
		IsPersistent:    dep.IsPersistent(),
		ValidPeriod:     dep.ValidPeriod(),
		IsRecordEnabled: dep.IsRecordEnabled(),
	}
}
