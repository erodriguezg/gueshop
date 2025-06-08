package app

import (
	_ "github.com/erodriguezg/gueshop/docs" // auto-generado por swag
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/fx"
)

type RoutesParams struct {
	fx.In

	App *fiber.App
}

// ProvideRoutes registra ruta de infraestructura
func ProvideInfrastructureRoutes(p RoutesParams) {

	// Ruta de salud
	p.App.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Swagger UI
	p.App.Get("/swagger/*", fiberSwagger.WrapHandler)
}
