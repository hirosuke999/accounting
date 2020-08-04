package handler

import (
	"accounting/src/infrastructure/store"
	uc "accounting/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllAccountTitle 勘定科目を作成する
func GetAllAccountTitle(ctx *gin.Context) {

	repo := store.NewAccountTitleStore()
	output := uc.GetAllAccountTitle(*repo)

	ctx.JSON(http.StatusOK, output)
}
