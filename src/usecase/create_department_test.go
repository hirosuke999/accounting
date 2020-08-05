package usecase

import (
	"accounting/src/domain/model"
	"accounting/src/domain/model/value"
	"accounting/src/util/test/mock_repository"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateDepartment(t *testing.T) {
	t.Run("部門を作成する", func(t *testing.T) {
		var (
			id              = uint(1)
			code            = "1000"
			name            = "開発部"
			isPersistent    = true
			periodFrom      = time.Now()
			periodTo        = time.Now()
			validPeriod     = *value.NewPeriod(&value.PeriodParam{From: periodFrom, To: periodTo})
			isRecordEnabled = true
		)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock_repository.NewMockDepartment(ctrl)
		repo.
			EXPECT().
			Save(gomock.Any()).
			Return(model.NewDepartment(&model.DepartmentParam{
				ID:              id,
				Code:            code,
				Name:            name,
				IsPersistent:    isPersistent,
				ValidPeriod:     validPeriod,
				IsRecordEnabled: isRecordEnabled,
			}))

		output := CreateDepartment(repo, &CreateDepartmentInput{
			Code:            code,
			Name:            name,
			IsPersistent:    isPersistent,
			ValidPeriod:     validPeriod,
			IsRecordEnabled: isRecordEnabled,
		})

		t.Log(output)
		assert.Equal(t, id, output.ID)
		assert.Equal(t, code, output.Code)
		assert.Equal(t, name, output.Name)
		assert.Equal(t, isPersistent, output.IsPersistent)
		assert.Equal(t, validPeriod, output.ValidPeriod)
		assert.Equal(t, isRecordEnabled, output.IsRecordEnabled)
	})
}
