package mysql

import "gorm.io/gorm"

type insightDatas struct {
	db *gorm.DB
}

func newInsightDatas(ds *datastore) *insightDatas {
	return &insightDatas{db: ds.db}
}
