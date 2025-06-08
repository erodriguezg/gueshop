package db

import (
	"fmt"

	"github.com/erodriguezg/gueshop/internal/util"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// NewMigrator configura el objeto Migrate para usar dentro de Fx
func NewMigrator(db *sqlx.DB, props util.ConfigProperties) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	migrationsPath := props.GetProp("MIGRATIONS_PATH")
	if migrationsPath == "" {
		migrationsPath = "internal/db/migrations"
	}

	sourceURL := fmt.Sprintf("file://%s", migrationsPath)

	m, err := migrate.NewWithDatabaseInstance(
		sourceURL,
		"postgres", driver,
	)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// RunMigrations se invoca desde fx.Invoke para aplicar migraciones al iniciar
func RunMigrations(m *migrate.Migrate, props util.ConfigProperties, logger *zap.Logger) error {

	mayRunMigrations := props.GetBoolProp("MIGRATIONS_ENABLED")

	if !mayRunMigrations {
		logger.Info("migrations are deactivated. not running migrations!")
	}

	err := m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("falló migración: %w", err)
	}

	logger.Info("migrations runs successfully!")

	return nil
}
