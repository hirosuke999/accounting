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

	db := store.GetDB()
	repos := store.NewAccountTitleStore(db)

	output := uc.CreateAccountTitle(repos, &input)

	ctx.JSON(http.StatusOK, output)
}
