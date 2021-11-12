package mysql

import "gorm.io/gorm"

type insightFields struct {
	db *gorm.DB
}

func newInsightFields(ds *datastore) *insightFields {
	return &insightFields{db: ds.db}
}