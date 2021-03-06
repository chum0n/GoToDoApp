package store

import (
	"strconv"
	"todo/repository"

	"github.com/jinzhu/gorm"
)

// Store storesテーブルデータ
type Store struct {
	gorm.Model
	Store_id   string
	Store_name string
	Address    string
	Price      int
}

func init() {
	// 自動マイグレーションはテーブルや不足しているカラムとインデックスのみ生成します。データ保護のため、既存のカラム型の変更や未使用のカラムの削除はしません。
	repository.DB.AutoMigrate(&Store{})
}

// Insert レコードを登録
func Insert(tx *gorm.DB, store_id string, store_name string, address string, priceS string) error {
	price, _ := strconv.Atoi(priceS)
	err := tx.Create(&Store{Store_id: store_id, Store_name: store_name, Address: address, Price: price}).Error
	return err
}

// SelectAllStores storesテーブルの全レコードを取得する
func SelectAllStores() []Store {
	var stores []Store
	// db.Findで構造体に対するテーブルの要素全てを取得し、それをOrder("created_at desc)で新しいものが上に来るよう並び替えを行なっています。
	repository.DB.Order("created_at desc").Find(&stores)
	return stores
}

// SelectByStoreID storeIDを条件にレコードを取得する
func SelectByStoreID(store_id string) Store {
	var store Store
	// 取得されるのは絶対一件なのでFirstを使った
	repository.DB.Where("store_id = ?", store_id).First(&store)
	return store
}

// UpdateByStoreID storeIDを条件にレコードを更新する
func UpdateByStoreID(store_id string, store_name string, address string, priceS string) {
	price, _ := strconv.Atoi(priceS)
	var store Store
	// 特定のレコードを呼び出す
	repository.DB.Where("store_id = ?", store_id).First(&store)
	store.Store_name = store_name
	store.Address = address
	store.Price = price
	repository.DB.Save(&store)
}

// DeleteByStoreID
func DeleteByStoreID(store_id string) {
	var store Store
	repository.DB.Where("store_id = ?", store_id).First(&store)
	repository.DB.Delete(&store)
}

// SearchByPrice
func SearchByPrice(lowerS string, upperS string) []Store {
	var stores []Store
	lower, _ := strconv.Atoi(lowerS)
	upper, _ := strconv.Atoi(upperS)
	repository.DB.Where("price >= ? AND price <= ?", lower, upper).Find(&stores)
	return stores
}

// SearchByAddress
func SearchByAddress(address string) []Store {
	var stores []Store
	repository.DB.Where("address = ?", address).Find(&stores)
	return stores
}

// SearchByPriceAndAddress
func SearchByPriceAndAddress(lowerS string, upperS string, address string) []Store {
	var stores []Store
	lower, _ := strconv.Atoi(lowerS)
	upper, _ := strconv.Atoi(upperS)
	if address == "" {
		repository.DB.Where("price >= ? AND price <= ?", lower, upper).Find(&stores)
	} else {
		repository.DB.Where("price >= ? AND price <= ? AND address = ?", lower, upper, address).Find(&stores)
	}
	return stores
}

// CountAllStore
func CountAllStore() int {
	// var stores []Store
	var count int
	// repository.DB.Where("deleted_at != ?", "NULL").Find(&stores).Count(&count)
	repository.DB.Table("stores").Count(&count)
	return count
}
