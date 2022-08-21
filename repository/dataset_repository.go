package repository

import (
	"data-engine/config"
	"gorm.io/gorm"
)

type DataSetRepository struct {
	conf.Database
	logger conf.Logger
}

func NewDataSetRepository(db conf.Database, logger conf.Logger) DataSetRepository {
	return DataSetRepository{
		Database: db,
		logger:   logger,
	}
}

func (r DataSetRepository) WithTrx(trxHandle *gorm.DB) DataSetRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
