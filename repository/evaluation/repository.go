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

type ForRanking struct {
	Store_name string
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

func Ranking() []ForRanking {
	var forRankings []ForRanking
	repository.DB.Table("stores").Select("stores.Store_name, evaluations.Evaluation").Joins("left join evaluations on stores.Store_id = evaluations.Store_id").Order("evaluation desc").Scan(&forRankings)
	return forRankings
}
