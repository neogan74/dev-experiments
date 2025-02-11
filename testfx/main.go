package main

import "go.uber.org/fx"

func main() {
	appOpts := fx.Options(
		fx.Provide(NewConfig),
		fx.Provide(NewLogger),
		fx.Provide(NewEchoHandler),
		fx.Provide(NewMux),
	)

	fx.New(appOpts).Run()
}
