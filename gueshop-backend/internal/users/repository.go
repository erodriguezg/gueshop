package users

import (
    "context"

    "github.com/jmoiron/sqlx"
)

type UserRepository struct {
    db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user User) error {
    _, err := r.db.NamedExecContext(ctx,
        `INSERT INTO users (name, email) VALUES (:name, :email)`, user)
    return err
}

func (r *UserRepository) List(ctx context.Context) ([]User, error) {
    var users []User
    err := r.db.SelectContext(ctx, &users, `SELECT id, name, email FROM users`)
    return users, err
}

