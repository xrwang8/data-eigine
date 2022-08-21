package routers

import (
	"data-engine/config"
	"data-engine/routers/controllers"
)

type ModelRoutes struct {
	logger          conf.Logger
	handler         conf.RequestHandler
	modelController controllers.ModelController
}

func (s ModelRoutes) Setup() {

	api := s.handler.Gin.Group("/api/v1")
	{
		api.POST("/model", s.modelController.SaveModel)
		api.DELETE("/model/:mid", s.modelController.DeleteModel)
		api.GET("/model/:mid", s.modelController.GetModel)
	}
}

func NewModelRoutes(logger conf.Logger, handler conf.RequestHandler,
	modelController controllers.ModelController) ModelRoutes {
	return ModelRoutes{
		handler:         handler,
		logger:          logger,
		modelController: modelController,
	}
}
