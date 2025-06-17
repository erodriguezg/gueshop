package catalog

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type CatalogRepository struct {
	db *sqlx.DB
}

func NewCatalogRepository(db *sqlx.DB) *CatalogRepository {
	return &CatalogRepository{db: db}
}

func (r *CatalogRepository) GetParentCategories(ctx context.Context) ([]Category, error) {
	var categories []Category
	query := `
		select id, name, parent_id, description 
		from cat_categories
		where parent_id is null
		order by name`
	err := r.db.SelectContext(ctx, &categories, query)
	return categories, err
}
