package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// var UserFactory = factory.NewFactory(
// 	&AccountTitle{},
// ).Attr("Name", func(args factory.Args) (interface{}, error) {
// 	accountTitle := args.Instance().(*AccountTitle)
// 	return fmt.Sprintf("user-%d", accountTitle), nil
// })

func TestConsumptionTaxType(t *testing.T) {
	t.Run("success NewConsumptionTaxType()", func(t *testing.T) {
		name := "課税売上げ10%"
		taxType := NewConsumptionTaxType(name)
		t.Log(taxType)
		assert.Equal(t, name, taxType.Name)
	})
}
