package model

import (
	"accounting/src/domain/model/value"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDepartment(t *testing.T) {
	t.Run("部門を生成する", func(t *testing.T) {
		var (
			code        = "1000"
			name        = "開発部"
			valuePeriod = *value.NewPeriod(&value.PeriodParam{From: time.Now(), To: time.Now()})
		)

		output := NewDepartment(&DepartmentParam{
			Code:        code,
			Name:        name,
			ValidPeriod: valuePeriod,
		})

		t.Log(output)
		assert.Equal(t, uint(0), output.ID())
		assert.Equal(t, code, output.Code())
		assert.Equal(t, name, output.Name())
		assert.Equal(t, false, output.IsPersistent())
		assert.Equal(t, valuePeriod, output.ValidPeriod())
		assert.Equal(t, false, output.IsRecordEnabled())
	})
}
