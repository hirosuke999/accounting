package usecase

import (
	"accounting/src/domain/model"
	"accounting/src/util/test/mock_repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateSalesTaxType(t *testing.T) {
	t.Run("消費税区分を作成する", func(t *testing.T) {
		var (
			id   = uint(1)
			name = "課税売上げ10%"
		)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock_repository.NewMockSalesTaxType(ctrl)
		repo.
			EXPECT().
			Save(gomock.Any()).
			Return(&model.SalesTaxType{ID: id, Name: name})

		output := CreateSalesTaxType(repo, CreateSalesTaxTypeInput{Name: name})

		// t.Log(output)
		assert.Equal(t, id, output.ID)
		assert.Equal(t, name, output.Name)
	})
}
