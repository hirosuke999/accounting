package enum

// ProjectNecessity プロジェクト要否
type ProjectNecessity int

const (
	// ProjectNecessityNone プロジェクト要否: 不要
	ProjectNecessityNone ProjectNecessity = iota + 1
	// ProjectNecessityOptional プロジェクト要否: 任意
	ProjectNecessityOptional
	// ProjectNecessityRequired プロジェクト要否: 必須
	ProjectNecessityRequired
)

func (s ProjectNecessity) String() string {
	return [...]string{"None", "Optional", "Required"}[s-1]
}
