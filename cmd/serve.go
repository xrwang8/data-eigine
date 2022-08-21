package cmd

import (
	"data-engine/config"
	"data-engine/routers"
	"github.com/spf13/cobra"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {

}

func (s *ServeCommand) Run() conf.CommandRunner {
	return func(
		conf conf.Config,
		router conf.RequestHandler,
		route routers.Routes,
		logger conf.Logger,
		database conf.Database,
	) {
		//middleware.Setup()
		route.Setup()

		logger.Info("Running server")
		if conf.App.Port == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + conf.App.Port)
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
