package main

import (
	"log"
	"todo/repository"
	"todo/repository/customer"
	"todo/repository/evaluation"
	"todo/repository/store"

	"github.com/gin-gonic/gin"

	// 使わないことを明示しているが必要
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := gin.Default()
	// 静的ファイル
	// router.Static("URL", "静的ファイル格納場所")
	router.Static("/assets", "./assets")
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
		stores := store.SelectAllStores()
		count := store.CountAllStore()
		ctx.HTML(200, "index.html", gin.H{
			"stores": stores,
			"count":  count,
		})
	})

	// signup
	router.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(200, "signup.html", gin.H{})
	})

	// signin
	router.GET("/signin", func(ctx *gin.Context) {
		ctx.HTML(200, "signin.html", gin.H{})
	})

	// searchstore
	router.GET("/searchstore", func(ctx *gin.Context) {
		stores := store.SelectAllStores()
		ctx.HTML(200, "searchstore.html", gin.H{
			"stores": stores,
		})
	})

	router.POST("/consearch", func(ctx *gin.Context) {
		lower := ctx.PostForm("lower")
		upper := ctx.PostForm("upper")
		address := ctx.PostForm("address")
		stores := store.SearchByPriceAndAddress(lower, upper, address)
		ctx.HTML(200, "searchstore.html", gin.H{
			"stores": stores,
		})
	})

	// editstore
	router.GET("/editstore", func(ctx *gin.Context) {
		stores := store.SelectAllStores()
		// evals := make([]evaluation.Evaluation, 0, len(stores))
		// for i := 0; i < len(stores); i++ {
		// 	evals[i] = evaluation.SelectEvaluation(stores[i].Store_id)
		// }
		ctx.HTML(200, "editstore.html", gin.H{
			"stores": stores,
		})
	})

	// ranking
	router.GET("/ranking", func(ctx *gin.Context) {
		stores := store.SelectAllStores()
		ctx.HTML(200, "ranking.html", gin.H{
			"stores": stores,
		})
	})

	//Create new customer
	router.POST("/new", func(ctx *gin.Context) {
		customer_id := ctx.PostForm("customer_id")
		customer_name := ctx.PostForm("customer_name")
		age := ctx.PostForm("age")
		gender := ctx.PostForm("gender")
		customer.Insert(customer_id, customer_name, age, gender)
		// localhost:8080/にステータスコード302としてリダイレクト
		ctx.Redirect(302, "/")
	})

	// Create new store
	router.POST("/newstore", func(ctx *gin.Context) {
		store_id := ctx.PostForm("store_id")
		store_name := ctx.PostForm("store_name")
		address := ctx.PostForm("address")
		price := ctx.PostForm("price")
		evaluationS := ctx.PostForm("evaluation")
		// transaction
		tx := repository.DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		if err := tx.Error; err != nil {
			return
		}

		if err := store.Insert(tx, store_id, store_name, address, price); err != nil {
			tx.Rollback()
			return
		}

		if err := evaluation.Insert(tx, store_id, evaluationS); err != nil {
			tx.Rollback()
			return
		}

		tx.Commit()

		// localhost:8080/にステータスコード302としてリダイレクト
		ctx.Redirect(302, "/editstore")
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

	//Detail new store
	router.GET("/storedetail/:store_id", func(ctx *gin.Context) {
		// n := ctx.Param("customer_id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic(err)
		// }
		store_id := ctx.Param("store_id")
		log.Println(store_id)
		store := store.SelectByStoreID(store_id)
		ctx.HTML(200, "storedetail.html", gin.H{"store": store})
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

	//Update store
	router.POST("/updatestore/:store_id", func(ctx *gin.Context) {
		// ここでidの値を受け取り、int型に変換
		// n := ctx.Param("id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic("ERROR")
		// }
		store_id := ctx.Param("store_id")
		store_name := ctx.PostForm("store_name")
		address := ctx.PostForm("address")
		price := ctx.PostForm("price")
		store.UpdateByStoreID(store_id, store_name, address, price)
		ctx.Redirect(302, "/editstore")
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

	// 店舗削除確認
	router.GET("/deletestore_check/:store_id", func(ctx *gin.Context) {
		// n := ctx.Param("id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic("ERROR")
		// }
		store_id := ctx.Param("store_id")
		store := store.SelectByStoreID(store_id)
		ctx.HTML(200, "deletestore.html", gin.H{"store": store})
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
		ctx.Redirect(302, "/editstore")
	})

	//Delete
	router.POST("/deletestore/:store_id", func(ctx *gin.Context) {
		// n := ctx.Param("id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic("ERROR")
		// }
		store_id := ctx.Param("store_id")
		store.DeleteByStoreID(store_id)
		ctx.Redirect(302, "/")
	})

	// search by price
	router.POST("/pricesearch", func(ctx *gin.Context) {
		// n := ctx.Param("id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic("ERROR")
		// }
		lower := ctx.PostForm("lower")
		upper := ctx.PostForm("upper")
		stores := store.SearchByPrice(lower, upper)
		ctx.HTML(200, "pricesearch.html", gin.H{"stores": stores})
	})

	// search by address
	router.POST("/addresssearch", func(ctx *gin.Context) {
		// n := ctx.Param("id")
		// id, err := strconv.Atoi(n)
		// if err != nil {
		// 	panic("ERROR")
		// }
		address := ctx.PostForm("address")
		stores := store.SearchByAddress(address)
		ctx.HTML(200, "addresssearch.html", gin.H{"stores": stores})
	})

	router.Run()
}
