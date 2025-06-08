package app

import (
	"context"
	"fmt"

	"github.com/erodriguezg/gueshop/internal/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ProvideFiber crea la instancia de Fiber con json acelerado
func ProvideFiber() *fiber.App {
	return fiber.New(fiber.Config{
		AppName:     "GueShop Application",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
}

// StartServer lanza el servidor
func StartServer(lc fx.Lifecycle, props util.ConfigProperties, app *fiber.App, logger *zap.Logger) {
	port := props.GetProp("FIBER_PORT")
	if port == "" {
		port = "3000"
	}

	addr := fmt.Sprintf(":%s", port)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Iniciando servidor HTTP", zap.String("addr", addr))
			go func() {
				if err := app.Listen(addr); err != nil {
					logger.Fatal("Error al iniciar el servidor", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Apagando servidor")
			return app.Shutdown()
		},
	})
}
