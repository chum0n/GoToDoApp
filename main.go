package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// 使わないことを明示しているが、いる
	_ "https://github.com/lib/pq"
)

// モデル設計
type Customer struct {
	// gorm.ModelはID, CreatedAt, UpdatedAt, DeletedAtというフィールドを持つ、GoのStructです。
	// あなたのモデルに組み込んで使っても良いですし、組み込まずに独自のモデルを使っても構いません。
	// `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`フィールドを`Customer`モデルに注入します
	// いつかタイムスタンプも使うかもしれないので追加
	gorm.Model
	Customer_id string
	Customer_name string
	Age int
	Gender int
}

//DB初期化
func dbInit() {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbInit）")
	}
	// 自動マイグレーションはテーブルや不足しているカラムとインデックスのみ生成します。データ保護のため、既存のカラム型の変更や未使用のカラムの削除はしません。
    db.AutoMigrate(&Customer{})
    defer db.Close()
}

//DB追加
func dbInsert(text string, status string) {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbInsert）")
	}
    db.Create(&Todo{Text: text, Status: status})
    defer db.Close()
}

//DB更新
func dbUpdate(id int, text string, status string) {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbUpdate）")
	}
    var todo Todo
    db.First(&todo, id)
    todo.Text = text
    todo.Status = status
    db.Save(&todo)
    db.Close()
}

//DB削除
func dbDelete(id int) {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbDelete）")
	}
    var todo Todo
    db.First(&todo, id)
    db.Delete(&todo)
    db.Close()
}

//DB全取得
func dbGetAll() []Todo {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbGetAll）")
	}
    var todos []Todo
    db.Order("created_at desc").Find(&todos)
    db.Close()
    return todos
}

//DB一つ取得
func dbGetOne(id int) Todo {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbGetOne）")
	}
    var todo Todo
    db.First(&todo, id)
    db.Close()
    return todo
}

func main() {
	router := gin.Default()
	// HTMLを読み込むディレクトリを指定
	router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin!!"

	// index.htmlにGETで繋ぐ
	router.GET("/", func(ctx *gin.Context) {
		// mapで値を渡す
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run()
}
