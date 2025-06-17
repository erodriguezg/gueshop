package catalog

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewCatalogRepository,
		NewCatalogService,
	),
	fx.Invoke(RegisterRoutes),
)
