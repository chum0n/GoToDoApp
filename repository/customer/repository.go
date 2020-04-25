package customer

import (
	"strconv"
	"todo/repository"

	"github.com/jinzhu/gorm"
)

// Customer　customersテーブルデータ
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
	Customer_id   string
	Customer_name string
	Age           int
	Gender        int
}

func init() {
	// 自動マイグレーションはテーブルや不足しているカラムとインデックスのみ生成します。データ保護のため、既存のカラム型の変更や未使用のカラムの削除はしません。
	repository.DB.AutoMigrate(&Customer{})
}

// Insert レコードを登録する
func Insert(customer_id string, customer_name string, ageS string, genderS string) {
	age, _ := strconv.Atoi(ageS)
	gender, _ := strconv.Atoi(genderS)
	// Customerという構造体に与えられた引数をいれた状態で、db.Create()に渡しています。
	repository.DB.Create(&Customer{Customer_id: customer_id, Customer_name: customer_name, Age: age, Gender: gender})
}

// SelectAllCustomers customersテーブルの全レコードを取得する
func SelectAllCustomers() []Customer {
	var customers []Customer
	// db.Find(&customers)で構造体Customerに対するテーブルの要素全てを取得し、それをOrder("created_at desc)で新しいものが上に来るよう並び替えを行なっています。
	repository.DB.Order("created_at desc").Find(&customers)
	return customers
}

// SelectByCustomerID customerIDを条件にレコードを取得する
func SelectByCustomerID(customer_id string) Customer {
	var customer Customer
	// 取得されるのは絶対一件なのでFirstを使った
	repository.DB.Where("customer_id = ?", customer_id).First(&customer)
	return customer
}

//Update
func UpdateByCustomerID(customer_id string, customer_name string, ageS string, genderS string) {
	age, _ := strconv.Atoi(ageS)
	gender, _ := strconv.Atoi(genderS)
	var customer Customer
	// 特定のレコードを呼び出す
	repository.DB.Where("customer_id = ?", customer_id).First(&customer)
	customer.Customer_name = customer_name
	customer.Age = age
	customer.Gender = gender
	repository.DB.Save(&customer)
}

//DB削除
func DeleteByCustomerID(customer_id string) {
	var customer Customer
	repository.DB.Where("customer_id = ?", customer_id).First(&customer)
	repository.DB.Delete(&customer)
}
