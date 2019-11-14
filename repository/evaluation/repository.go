package evaluation

import (
	"strconv"
	"todo/repository"

	"github.com/jinzhu/gorm"
)

type Evaluation struct {
	Store_id   string
	Evaluation int
}

func init() {
	repository.DB.AutoMigrate(&Evaluation{})
}

// Insert
func Insert(tx *gorm.DB, store_id string, evaluationS string) error {
	evaluation, _ := strconv.Atoi(evaluationS)
	err := tx.Create(&Evaluation{Store_id: store_id, Evaluation: evaluation}).Error
	return err
}

func SelectEvaluation(store_id string) Evaluation {
	var eval Evaluation
	repository.DB.Where("store_id = ?", store_id).Find(&eval)
	return eval
}
