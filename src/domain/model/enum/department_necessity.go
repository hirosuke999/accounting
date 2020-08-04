package enum

// DepartmentNecessity 部門要否
type DepartmentNecessity int

const (
	// DepartmentNecessityNone 部門要否: 不要
	DepartmentNecessityNone DepartmentNecessity = iota + 1
	// DepartmentNecessityOptional 部門要否: 任意
	DepartmentNecessityOptional
	// DepartmentNecessityRequired 部門要否: 必須
	DepartmentNecessityRequired
)

func (s DepartmentNecessity) String() string {
	return [...]string{"None", "Optional", "Required"}[s-1]
}
