package controllers

import (
	conf "data-engine/config"
	"data-engine/http_dto"
	"data-engine/models"
	"data-engine/pkg/app"
	"data-engine/pkg/e"
	"data-engine/services"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"math/rand"
	"net/http"
)

type ModelController struct {
	service services.ModelService
	logger  conf.Logger
}

func NewDataModelController(modelService services.ModelService, logger conf.Logger) ModelController {
	return ModelController{
		service: modelService,
		logger:  logger,
	}
}

// SaveModel  创建模型
// @Tags 模型管理
// @Summary 创建模型
// @Description 获取JSON
// @Accept  application/json
// @Product application/json
// @Param data body http_dto.ModelCreateReq true "data"
// @Success 200 {object} app.Response
// @Router /api/v1/model [post]
func (m *ModelController) SaveModel(c *gin.Context) {
	appG := app.Gin{C: c}
	modelReq := http_dto.ModelCreateReq{}
	if err := c.ShouldBindJSON(&modelReq); err != nil {
		m.logger.Error(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err)
		return
	}

	// todo 分中心适配器接口，一下dataset模拟返回结果
	model := models.Model{
		Name:         modelReq.Name,
		CenterId:     modelReq.CenterId,
		UserId:       modelReq.UserId,
		NonPublic:    modelReq.NonPublic,
		ProviderType: modelReq.ProviderType,
		Mid:          modelReq.Mid,
		Sid:          rand.Intn(1000),
	}

	if err := m.service.CreateModel(model); err != nil {
		m.logger.Error(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err)
		return
	}

	resp := &http_dto.ModelResp{
		Name:         model.Name,
		NonPublic:    model.NonPublic,
		ProviderType: model.ProviderType,
		Mid:          model.Mid,
		Sid:          model.Sid,
	}
	appG.Response(http.StatusOK, e.SUCCESS, resp)

}

// DeleteModel 删除模型
// @Summary 删除模型
// @Description 获取JSON
// @Tags 模型管理
// @Param id path string true "mid"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/model/{mid} [delete]
func (d *ModelController) DeleteModel(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("mid")).MustInt()
	if err := d.service.DeleteModel(uint(id)); err != nil {
		d.logger.Error(err)
		appG.Response(http.StatusBadRequest, e.ERROR_DELETE_DATASET_FAIL, err)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// GetModel  根据主中心查询分中心模型
// @Summary 查询分中心模型
// @Description 获取JSON
// @Tags 模型管理
// @Param id path string true "mid"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/model/{mid} [get]
func (d *ModelController) GetModel(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("mid")).MustInt()
	modelList, err := d.service.GetModel(id)
	if err != nil {
		d.logger.Error(err)
		appG.Response(http.StatusBadRequest, e.ERROR_SELECT_DATASET_FAIL, err)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, modelList)

}
