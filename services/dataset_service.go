package services

import (
	"data-engine/config"
	"data-engine/http_dto"
	"data-engine/models"
	"data-engine/repository"
)

type DataSetService struct {
	logger     conf.Logger
	repository repository.DataSetRepository
}

func NewDataSetService(logger conf.Logger, repository repository.DataSetRepository) DataSetService {
	return DataSetService{
		logger:     logger,
		repository: repository,
	}
}

func (s DataSetService) CreateDataSet(dataset models.DataSet) error {
	return s.repository.Create(&dataset).Error
}

func (s DataSetService) DeleteDataSet(mid uint) error {
	return s.repository.Where("mid=?", mid).Delete(&models.DataSet{}).Error
}

func (s DataSetService) GetDataSet(mid int) (datasets []http_dto.DatasetCreateResp, err error) {
	return datasets, s.repository.Model(&models.DataSet{}).Where("mid=?", mid).Find(&datasets).Error
}
