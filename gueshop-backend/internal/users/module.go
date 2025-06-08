package users

import "go.uber.org/fx"

// Module para fx.App
var Module = fx.Options(
    fx.Provide(
        NewUserRepository,
        NewUserService,
    ),
    fx.Invoke(RegisterRoutes),
)

