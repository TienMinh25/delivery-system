package main

import (
	"github.com/TienMinh25/delivery-system/internal/env"
	"go.uber.org/fx"
	"log"
)

func main() {
	app := fx.New(
		fx.Provide(env.NewEnvManager),
		fx.Invoke(func(env *env.EnvManager) {
			log.Println("✅ Loaded environment variables successfully")
		}))

	app.Run()
}
