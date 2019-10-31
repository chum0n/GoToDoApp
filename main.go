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
	// 	GORMの標準のモデルはどういう名前で入っているかというと、
	// id → ID
	// created_at → CreatedAt
	// updated_at → UpdatedAt
	// deleted_at → DeletedAt
	// となっています。これらはHTMLにGO側から変数を渡した時の呼び出すときにも使うので注意してください。
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
func dbInsert(customer_id string, customer_name string, age int, gender int) {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbInsert）")
	}
	// Customerという構造体に与えられた引数をいれた状態で、db.Create()に渡しています。
    db.Create(&Customer{Customer_id: customer_id, Customer_name: customer_name, Age: age, Gender: gender})
    defer db.Close()
}

//DB更新
func dbUpdate(customer_id string, customer_name string, age int, gender int) {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbUpdate）")
	}
	var customer Customer
	// 特定のレコードを呼び出す
    db.First(&customer, customer_id)
	customer.Customer_name = customer_name
	customer.Age = age
    customer.Gender = gender
    db.Save(&customer)
    db.Close()
}

//DB削除
func dbDelete(id int) {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbDelete）")
	}
    var customer Customer
    db.First(&customer, customer_id)
    db.Delete(&customer)
    db.Close()
}

//DB全取得
func dbGetAll() []Customer {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbGetAll）")
	}
	var customers []Customer
	// db.Find(&customers)で構造体Customerに対するテーブルの要素全てを取得し、それをOrder("created_at desc)で新しいものが上に来るよう並び替えを行なっています。
    db.Order("created_at desc").Find(&customers)
    db.Close()
    return customers
}

//DB一つ取得
func dbGetOne(customer_id string) Customer {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=daisuke dbname=ex4 password=")
    if err != nil {
        panic("データベース開けず！（dbGetOne）")
	}
	var customer Customer
	// 第２引数にはidを加えることで特定のレコードを取得することができます。
    db.First(&customer, customer_id)
    db.Close()
    return customer
}

func main() {
	router := gin.Default()
	// HTMLを読み込むディレクトリを指定
	router.LoadHTMLGlob("templates/*.html")

	// data := "Hello Go/Gin!!"

	// // index.htmlにGETで繋ぐ
	// router.GET("/", func(ctx *gin.Context) {
	// 	// mapで値を渡す
	// 	ctx.HTML(200, "index.html", gin.H{"data": data})
	// })

	dbInit()

	//Index
    router.GET("/", func(ctx *gin.Context) {
        todos := dbGetAll()
        ctx.HTML(200, "index.html", gin.H{
            "customers": customers,
        })
	})

	//Create
    router.POST("/new", func(ctx *gin.Context) {
		customer_id := ctx.PostForm("customer_id")
		customer_name := ctx.PostForm("customer_name")
		age := ctx.PostForm("age")
		gender := ctx.PostForm("gender")
		dbInsert(customer_id, customer_name, age, gender)
		// localhost:8080/にステータスコード302としてリダイレクト
        ctx.Redirect(302, "/")
	})

	//Detail
    router.GET("/detail/:customer_id", func(ctx *gin.Context) {
        // n := ctx.Param("customer_id")
        // id, err := strconv.Atoi(n)
        if err != nil {
            panic(err)
        }
        todo := dbGetOne(customer_id)
        ctx.HTML(200, "detail.html", gin.H{"customer": customer})
	})
	
	//Update
    router.POST("/update/:customer_id", func(ctx *gin.Context) {
        // n := ctx.Param("id")
        // id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
		}
		customer_name := ctx.PostForm("customer_name")
        age := ctx.PostForm("age")
        gender := ctx.PostForm("gender")
        dbUpdate(customer_id, customer_name, age, gender)
        ctx.Redirect(302, "/")
	})
	
	//削除確認
    router.GET("/delete_check/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        todo := dbGetOne(id)
        ctx.HTML(200, "delete.html", gin.H{"todo": todo})
    })

    //Delete
    router.POST("/delete/:id", func(ctx *gin.Context) {
        n := ctx.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            panic("ERROR")
        }
        dbDelete(id)
        ctx.Redirect(302, "/")

    })

	router.Run()
}
