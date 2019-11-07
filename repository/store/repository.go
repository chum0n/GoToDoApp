package store

import (
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
func Insert(store_id string, store_name string, address string, price int) {
	repository.DB.Create(&Store{Store_id: store_id, Store_name: store_name, Address: address, Price: price})
}

// SelectAllStores storesテーブルの全レコードを取得する
func SelectAllStores() []Store {
	var stores []Store
	// db.Findで構造体に対するテーブルの要素全てを取得し、それをOrder("created_at desc)で新しいものが上に来るよう並び替えを行なっています。
	repository.DB.Order("created_at desc").Find(&stores)
	return stores
}
