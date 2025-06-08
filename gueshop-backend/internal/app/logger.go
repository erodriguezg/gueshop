package app

import (
	"github.com/erodriguezg/gueshop/internal/util"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// ProvideLogger lo usas en fx.Provide()
var ProvideLogger = fx.Provide(NewLogger)

// NewLogger retorna un *zap.Logger según APP_ENV
func NewLogger(props util.ConfigProperties) (*zap.Logger, error) {
	env := props.GetProp("APP_ENV")
	if env == "DEV" {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}
