package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"accounting/src/infrastructure/store"
	uc "accounting/src/usecase"
)

// CreateSalesTaxType 消費税区分を作成する
func CreateSalesTaxType(ctx *gin.Context) {
	var input uc.CreateSalesTaxTypeInput
	ctx.BindJSON(&input)

	repo := store.NewSalesTaxTypeStore()
	output := uc.CreateSalesTaxType(*repo, input)

	ctx.JSON(http.StatusOK, output)
}
