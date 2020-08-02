package handler

import (
	"accounting/src/infrastructure/store"
	uc "accounting/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllAccountTitle 勘定科目を作成する
func GetAllAccountTitle(ctx *gin.Context) {
	db := store.GetDB()
	repos := store.NewAccountTitleStore(db)

	output := uc.GetAllAccountTitle(repos)

	ctx.JSON(http.StatusOK, output)
}
