package enum

// BusinessPartnerNecessity 取引先要否
type BusinessPartnerNecessity int

const (
	// BusinessPartnerNecessityNone 取引先要否: 不要
	BusinessPartnerNecessityNone BusinessPartnerNecessity = iota + 1
	// BusinessPartnerNecessityOptional 取引先要否: 任意
	BusinessPartnerNecessityOptional
	// BusinessPartnerNecessityRequired 取引先要否: 必須
	BusinessPartnerNecessityRequired
)

func (s BusinessPartnerNecessity) String() string {
	return [...]string{"None", "Optional", "Required"}[s-1]
}
