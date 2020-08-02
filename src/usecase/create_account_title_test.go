package usecase

import (
	"testing"

	mock "accounting/src/util/test/mock_repository"

	m "accounting/src/domain/model"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountTitle(t *testing.T) {
	t.Run("勘定科目を作成する", func(t *testing.T) {
		var id uint = 777
		code := "4107"
		name := "売上高"
		// taxType := &ConsumptionTaxType{Name: "課税売上げ10%"}

		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		repos := mock.NewMockAccountTitleRepository(ctrl)
		repos.
			EXPECT().
			Save(gomock.Any()).
			Return(&m.AccountTitle{
				ID:   id,
				Code: code,
				Name: name,
			})

		output := CreateAccountTitle(repos, &CreateAccountTitleInput{Code: code, Name: name})

		t.Log(output)
		assert.Equal(t, id, output.ID)
		assert.Equal(t, code, output.Code)
		assert.Equal(t, name, output.Name)
		// assert.Equal(t, *taxType, accountTitle.TaxType)
	})
}
