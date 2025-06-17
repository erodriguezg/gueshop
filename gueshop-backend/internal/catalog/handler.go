package catalog

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registra las rutas del catálogo.
func RegisterRoutes(app *fiber.App, service *CatalogService) {
	group := app.Group("/api/catalog")
	group.Get("/categories", GetParentCategories(service))
}

// GetParentCategories maneja GET /api/catalog/categories
//
// @Summary Obtener categorías padre
// @Description Retorna todas las categorías sin padre (top-level)
// @Tags Categorías
// @Produce json
// @Success 200 {array} Category
// @Failure 500 {object} string
// @Router /api/catalog/categories [get]
func GetParentCategories(service *CatalogService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		categories, err := service.repo.GetParentCategories(context.Background())
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(categories)
	}
}
