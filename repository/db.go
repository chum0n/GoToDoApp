package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB 各repositoryで利用するDB接続情報
var DB *gorm.DB

// DBへの接続
func init() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=daisuke sslmode=disable")
	if err != nil {
		// panic("データベース開けず！（Init）")
		log.Println(err)
	}
}
