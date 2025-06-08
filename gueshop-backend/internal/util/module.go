package util

import (
	"go.uber.org/fx"
)

// Module de utilitarios base
var Module = fx.Options(
	// Proveedores de dependencias
	fx.Provide(
		NewGoEnvConfigProperties,
	),
)
