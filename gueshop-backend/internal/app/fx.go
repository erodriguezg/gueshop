package app

import (
	"go.uber.org/fx"

	"github.com/erodriguezg/gueshop/internal/db"
	"github.com/erodriguezg/gueshop/internal/users"
	"github.com/erodriguezg/gueshop/internal/util"
)

// Module representa el módulo raíz de la app, que ensambla todo
var Modules = fx.Options(

	// utilitarios transversales
	util.Module,

	// Logger
	ProvideLogger,

	// Template Renderer (Pongo2)
	ProvideTemplateRenderer,

	// DB + migraciones
	db.Module,

	// HTTP server y router
	fx.Provide(
		ProvideFiber,
	),
	fx.Invoke(
		ProvideInfrastructureRoutes,
		StartServer),

	// Módulos funcionales
	users.Module,
)
