package routers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewSwaggerRoutes),
	fx.Provide(NewDataSetRoutes),
	fx.Provide(NewModelRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(dataSetRoutes DataSetRoutes, swagger SwaggerRoutes, modelRoutes ModelRoutes) Routes {
	return Routes{
		dataSetRoutes,
		swagger,
		modelRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
