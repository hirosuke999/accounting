package model

// ConsumptionTaxType 消費税区分
type ConsumptionTaxType struct {
	ID   int
	Name string
}

// NewConsumptionTaxType 消費税区分を作成
func NewConsumptionTaxType(name string) *ConsumptionTaxType {
	return &ConsumptionTaxType{Name: name}
}
