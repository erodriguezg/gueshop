package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erodriguezg/gueshop/internal/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewDB crea y retorna el pool de conexión usando sqlx
func NewDB(props util.ConfigProperties) (*sqlx.DB, error) {
	dsn := props.GetProp("POSTGRES_URL")
	if dsn == "" {
		return nil, fmt.Errorf("POSTGRES_URL no está definido")
	}

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Opcional: verificar conexión al arrancar
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// WithTransaction permite ejecutar una función dentro de una transacción
func WithTransaction(ctx context.Context, db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
