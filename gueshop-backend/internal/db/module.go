package db

import (
    "go.uber.org/fx"
)

// Module de base de datos para Fx
var Module = fx.Options(
    // Proveedores de dependencias
    fx.Provide(
        NewDB,        // db.go → conexión con sqlx
        NewMigrator,  // migrate.go → instancia de migrate
    ),
    // Ejecutar migraciones al iniciar
    fx.Invoke(
        RunMigrations, // migrate.go
    ),
)

