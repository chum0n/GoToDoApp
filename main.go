package main

import (
	"log"
	"todo/repository"
	"todo/repository/customer"

	"github.com/gin-gonic/gin"

	// 使わないことを明示しているが、いる
	// _ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

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

	defer repository.DB.Close()

	//Index
	router.GET("/", func(ctx *gin.Context) {
		customers := customer.SelectAllCustomers()
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
		customer.Insert(customer_id, customer_name, age, gender)
		// localhost:8080/にステータスコード302としてリダイレクト
		ctx.Redirect(302, "/")
	})

	//Detail
	router.GET("/detail/:customer_id", func(ctx *gin.Context) {
		// n := ctx.Param("customer_id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic(err)
		// }
		customer_id := ctx.Param("customer_id")
		log.Println(customer_id)
		customer := customer.SelectByCustomerID(customer_id)
		ctx.HTML(200, "detail.html", gin.H{"customer": customer})
	})

	//Update
	router.POST("/update/:customer_id", func(ctx *gin.Context) {
		// ここでidの値を受け取り、int型に変換
		// n := ctx.Param("id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic("ERROR")
		// }
		customer_id := ctx.Param("customer_id")
		customer_name := ctx.PostForm("customer_name")
		age := ctx.PostForm("age")
		gender := ctx.PostForm("gender")
		customer.UpdateByCustomerID(customer_id, customer_name, age, gender)
		ctx.Redirect(302, "/")
	})

	//削除確認
	router.GET("/delete_check/:customer_id", func(ctx *gin.Context) {
		// n := ctx.Param("id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic("ERROR")
		// }
		customer_id := ctx.Param("customer_id")
		customer := customer.SelectByCustomerID(customer_id)
		ctx.HTML(200, "delete.html", gin.H{"customer": customer})
	})

	//Delete
	router.POST("/delete/:customer_id", func(ctx *gin.Context) {
		// n := ctx.Param("id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic("ERROR")
		// }
		customer_id := ctx.Param("customer_id")
		customer.DeleteByCustomerID(customer_id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
