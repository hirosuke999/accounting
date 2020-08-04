package usecase

import (
	"testing"

	. "accounting/src/domain/model"
	"accounting/src/domain/model/enum"
	"accounting/src/util/test/mock_repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountTitle(t *testing.T) {
	t.Run("勘定科目を作成する", func(t *testing.T) {
		var (
			id       = uint(777)
			code     = "4107"
			name     = "売上高"
			taxTypes = SalesTaxTypes{
				{ID: uint(1), Name: "課税売上げ8%"},
				{ID: uint(2), Name: "課税売上げ10%"},
				{ID: uint(3), Name: "非課税売上げ"},
			}
			departmentNecessity      = enum.DepartmentNecessityOptional
			segmentNecessity         = enum.SegmentNecessityOptional
			projectNecessity         = enum.ProjectNecessityOptional
			businessPartnerNecessity = enum.BusinessPartnerNecessityOptional
			reasonNoteNecessity      = enum.ReasonNoteNecessityOptional
		)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		repo := mock_repository.NewMockAccountTitle(ctrl)
		repo.
			EXPECT().
			GetSalesTaxTypes(gomock.Any()).
			Return(taxTypes)
		repo.
			EXPECT().
			Save(gomock.Any()).
			Return(&AccountTitle{
				ID:                       id,
				Code:                     code,
				Name:                     name,
				TaxTypes:                 taxTypes,
				DepartmentNecessity:      departmentNecessity,
				SegmentNecessity:         segmentNecessity,
				ProjectNecessity:         projectNecessity,
				BusinessPartnerNecessity: businessPartnerNecessity,
				ReasonNoteNecessity:      reasonNoteNecessity,
			})

		output := CreateAccountTitle(repo, CreateAccountTitleInput{
			Code:                     code,
			Name:                     name,
			TaxTypeIds:               []uint{1, 2, 3},
			DepartmentNecessity:      departmentNecessity,
			SegmentNecessity:         segmentNecessity,
			ProjectNecessity:         projectNecessity,
			BusinessPartnerNecessity: businessPartnerNecessity,
			ReasonNoteNecessity:      reasonNoteNecessity,
		})

		t.Log(output)
		assert.Equal(t, id, output.ID)
		assert.Equal(t, code, output.Code)
		assert.Equal(t, name, output.Name)
		assert.Equal(t, taxTypes, output.TaxTypes)
		assert.Equal(t, departmentNecessity, output.DepartmentNecessity)
		assert.Equal(t, segmentNecessity, output.SegmentNecessity)
		assert.Equal(t, projectNecessity, output.ProjectNecessity)
		assert.Equal(t, businessPartnerNecessity, output.BusinessPartnerNecessity)
		assert.Equal(t, reasonNoteNecessity, output.ReasonNoteNecessity)
	})
}
