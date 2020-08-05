package model

import "accounting/src/domain/model/value"

// Department 部門
type Department struct {
	id              ID
	code            string
	name            string
	isPersistent    bool
	validPeriod     value.Period
	isRecordEnabled bool
}

// DepartmentParam 部門生成パラメータ
type DepartmentParam struct {
	ID              ID
	Code            string
	Name            string
	IsPersistent    bool
	ValidPeriod     value.Period
	IsRecordEnabled bool
}

// NewDepartment 部門を生成する
func NewDepartment(param *DepartmentParam) *Department {
	validPeriod := param.ValidPeriod

	if value.Period.IsZero(validPeriod) {
		validPeriod = *value.NewPeriod(&value.PeriodParam{})
	}

	return &Department{
		id:              param.ID,
		code:            param.Code,
		name:            param.Name,
		isPersistent:    param.IsPersistent,
		validPeriod:     validPeriod,
		isRecordEnabled: param.IsRecordEnabled,
	}
}

// ID getter
func (d *Department) ID() ID {
	return d.id
}

// Code getter
func (d *Department) Code() string {
	return d.code
}

// Name getter
func (d *Department) Name() string {
	return d.name
}

// IsPersistent getter
func (d *Department) IsPersistent() bool {
	return d.isPersistent
}

// ValidPeriod getter
func (d *Department) ValidPeriod() value.Period {
	return d.validPeriod
}

// IsRecordEnabled getter
func (d *Department) IsRecordEnabled() bool {
	return d.isRecordEnabled
}
