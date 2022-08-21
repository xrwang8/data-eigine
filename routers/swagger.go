package routers

import (
	conf "data-engine/config"
	_ "data-engine/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerRoutes struct {
	handler conf.RequestHandler
}

func (s SwaggerRoutes) Setup() {
	s.handler.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func NewSwaggerRoutes(handler conf.RequestHandler) SwaggerRoutes {
	return SwaggerRoutes{
		handler: handler,
	}
}
