package store

import (
	"github.com/jinzhu/gorm"
	"../schema"
)

type EvaluationStore struct {
	*gorm.DB
}

func NewEvaluationStore(db *gorm.DB) *EvaluationStore {
	return &EvaluationStore{db}
}

func (d *EvaluationStore) CreateEvaluation(evaluation *schema.Evaluation) error {
	return d.Where(evaluation).FirstOrCreate(&evaluation).Error
}
