package services

import (
	"data-engine/config"
	"data-engine/http_dto"
	"data-engine/models"
	"data-engine/repository"
)

type ModelService struct {
	logger     conf.Logger
	repository repository.ModelRepository
}

func NewModelService(logger conf.Logger, repository repository.ModelRepository) ModelService {
	return ModelService{
		logger:     logger,
		repository: repository,
	}
}

func (s ModelService) CreateModel(model models.Model) error {
	return s.repository.Create(&model).Error
}

func (s ModelService) DeleteModel(mid uint) error {
	return s.repository.Where("mid=?", mid).Delete(&models.Model{}).Error
}

func (s ModelService) GetModel(mid int) (modelList []http_dto.ModelResp, err error) {
	return modelList, s.repository.Model(&models.Model{}).Where("mid=?", mid).Find(&modelList).Error
}
