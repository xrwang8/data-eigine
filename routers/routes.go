package routers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewDataSetRoutes),
	fx.Provide(NewSwaggerRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(dataSetRoutes DataSetRoutes, swagger SwaggerRoutes) Routes {
	return Routes{
		dataSetRoutes,
		swagger,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
