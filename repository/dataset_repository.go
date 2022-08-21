package repository

import (
	"data-engine/config"
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
