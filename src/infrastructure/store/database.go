package store

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // dependency package for gorm
)

// type Database struct {
// 	*gorm.DB
// }

// DB DB接続保持変数
var DB *gorm.DB

// Init DB初期化処理
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("db err: ", err)
		panic("データベースへの接続に失敗しました")
	}
	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

// GetDB DB参照を取得
func GetDB() *gorm.DB {
	return DB
}
