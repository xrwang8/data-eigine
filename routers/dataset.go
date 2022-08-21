package routers

import (
	"data-engine/config"
	"data-engine/routers/controllers"
)

type DataSetRoutes struct {
	logger            conf.Logger
	handler           conf.RequestHandler
	dataSetController controllers.DataSetController
}

func (s DataSetRoutes) Setup() {

	api := s.handler.Gin.Group("/api/v1")
	{
		api.POST("/dataset", s.dataSetController.SaveDataSet)
		api.DELETE("/dataset/:mid", s.dataSetController.DeleteDataSet)
		api.GET("/dataset/:mid", s.dataSetController.GetDataSet)
	}
}

func NewDataSetRoutes(logger conf.Logger, handler conf.RequestHandler,
	dataSetController controllers.DataSetController) DataSetRoutes {
	return DataSetRoutes{
		handler:           handler,
		logger:            logger,
		dataSetController: dataSetController,
	}
}
