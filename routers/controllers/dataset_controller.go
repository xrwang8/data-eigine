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

type DataSetController struct {
	service services.DataSetService
	logger  conf.Logger
}

func NewDataSetController(datasetService services.DataSetService, logger conf.Logger) DataSetController {
	return DataSetController{
		service: datasetService,
		logger:  logger,
	}
}

// SaveDataSet  创建数据集
// @Tags 数据集管理
// @Summary 创建数据集
// @Description 获取JSON
// @Accept  application/json
// @Product application/json
// @Param data body http_dto.DatasetCreateReq true "data"
// @Success 200 {object} app.Response
// @Router /api/v1/dataset [post]
func (d DataSetController) SaveDataSet(c *gin.Context) {
	appG := app.Gin{C: c}
	dataSetReq := http_dto.DatasetCreateReq{}
	if err := c.ShouldBindJSON(&dataSetReq); err != nil {
		d.logger.Error(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err)
		return
	}

	// todo 分中心适配器接口，一下dataset模拟返回结果
	dataSet := models.DataSet{
		Name:         dataSetReq.Name,
		CenterId:     dataSetReq.CenterId,
		UserId:       dataSetReq.UserId,
		NonPublic:    dataSetReq.NonPublic,
		ProviderType: dataSetReq.ProviderType,
		Mid:          dataSetReq.Mid,
		Sid:          rand.Intn(1000),
	}

	if err := d.service.CreateDataSet(dataSet); err != nil {
		d.logger.Error(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err)
		return
	}

	resp := &http_dto.DatasetCreateResp{
		Name:         dataSetReq.Name,
		NonPublic:    dataSetReq.NonPublic,
		ProviderType: dataSetReq.ProviderType,
		Mid:          dataSetReq.Mid,
		Sid:          dataSet.Sid,
	}
	appG.Response(http.StatusOK, e.SUCCESS, resp)

}

// DeleteDataSet 删除数据集
// @Summary 删除数据集
// @Description 获取JSON
// @Tags 数据集管理
// @Param id path string true "mid"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/dataset/{mid} [delete]
func (d DataSetController) DeleteDataSet(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("mid")).MustInt()
	if err := d.service.DeleteDataSet(uint(id)); err != nil {
		d.logger.Error(err)
		appG.Response(http.StatusBadRequest, e.ERROR_DELETE_DATASET_FAIL, err)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func (d DataSetController) GetDataSet(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("mid")).MustInt()
	datasets, err := d.service.GetDataSet(id)
	if err != nil {
		d.logger.Error(err)
		appG.Response(http.StatusBadRequest, e.ERROR_SELECT_DATASET_FAIL, err)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, datasets)

}
