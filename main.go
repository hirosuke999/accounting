package main

import (
	h "accounting/src/handler"
	"accounting/src/infrastructure/store"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // dependency package for gorm
)

// Migrate オートマイグレーション
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&store.AccountTitle{})
	db.AutoMigrate(&store.SalesTaxType{})
}

func main() {

	db := store.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")

	accountTitles := v1.Group("/account_titles")
	accountTitles.POST("/", h.CreateAccountTitle)
	accountTitles.GET("/", h.GetAllAccountTitle)

	salesTaxTypes := v1.Group("/sales_tax_types")
	salesTaxTypes.POST("/", h.CreateSalesTaxType)

	// testAuth := r.Group("/api/ping")
	// testAuth.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// Listen and Server in 0.0.0.0:8080
	r.Run("localhost:8080")
}
