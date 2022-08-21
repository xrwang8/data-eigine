package conf

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewConf),
	fx.Provide(GetLogger),
	fx.Provide(NewDatabase),
)
