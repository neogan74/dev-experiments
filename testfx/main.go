package main

import "go.uber.org/fx"

func main() {
	appOpts := fx.Options(
		fx.Provide(NewConfig),
	)

	fx.New(appOpts).Run()
}
