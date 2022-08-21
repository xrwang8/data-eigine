package cmd

import (
	"data-engine/config"
	"data-engine/repository"
	"data-engine/routers"
	"data-engine/routers/controllers"
	"data-engine/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routers.Module,
	conf.Module,
	services.Module,
	repository.Module,
)
