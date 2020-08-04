package enum

// ReasonNoteNecessity 増減理由要否
type ReasonNoteNecessity int

const (
	// ReasonNoteNecessityNone 増減理由要否: 不要
	ReasonNoteNecessityNone ReasonNoteNecessity = iota + 1
	// ReasonNoteNecessityOptional 増減理由要否: 任意
	ReasonNoteNecessityOptional
	// ReasonNoteNecessityRequired 増減理由要否: 必須
	ReasonNoteNecessityRequired
)

func (s ReasonNoteNecessity) String() string {
	return [...]string{"None", "Optional", "Required"}[s-1]
}
