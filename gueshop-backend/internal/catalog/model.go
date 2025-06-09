package catalog

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty" db:"parent_id"`
	Description string     `json:"description" db:"description"`
}

type Product struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	StockQty    int       `json:"stock_quantity" db:"stock_quantity"`
	CategoryID  uuid.UUID `json:"category_id" db:"category_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
