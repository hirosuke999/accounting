package value

import (
	t "time"
)

// Period 期間
type Period struct {
	from t.Time
	to   t.Time
}

// PeriodParam 生成パラメータ
type PeriodParam struct {
	From t.Time
	To   t.Time
}

// NewPeriod 期間を作成
func NewPeriod(param *PeriodParam) *Period {
	var (
		from = param.From
		to   = param.To
	)

	if t.Time.IsZero(from) {
		from = t.Now()
	}

	if t.Time.IsZero(to) {
		to = t.Now()
	}

	return &Period{from, to}
}

// IsZero ゼロ値チェック
func (p Period) IsZero() bool {
	return p == Period{}
}
