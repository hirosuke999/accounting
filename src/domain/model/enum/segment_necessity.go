package enum

// SegmentNecessity 事業セグメント要否
type SegmentNecessity int

const (
	// SegmentNecessityNone 事業セグメント要否: 不要
	SegmentNecessityNone SegmentNecessity = iota + 1
	// SegmentNecessityOptional 事業セグメント要否: 任意
	SegmentNecessityOptional
	// SegmentNecessityRequired 事業セグメント要否: 必須
	SegmentNecessityRequired
)

func (s SegmentNecessity) String() string {
	return [...]string{"None", "Optional", "Required"}[s-1]
}
