package repository

import (
	"data-engine/config"
)

type ModelRepository struct {
	conf.Database
	logger conf.Logger
}

func NewModelRepository(db conf.Database, logger conf.Logger) ModelRepository {
	return ModelRepository{
		Database: db,
		logger:   logger,
	}
}
