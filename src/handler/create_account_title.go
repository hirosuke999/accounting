package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"accounting/src/infrastructure/store"
	uc "accounting/src/usecase"
)

// CreateAccountTitle 勘定科目を作成する
func CreateAccountTitle(ctx *gin.Context) {
	var input uc.CreateAccountTitleInput
	ctx.BindJSON(&input)

	repo := store.NewAccountTitleStore()
	output := uc.CreateAccountTitle(*repo, input)

	ctx.JSON(http.StatusOK, output)
}
