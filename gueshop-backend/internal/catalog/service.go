package catalog

import "context"

type CatalogService struct {
	repo *CatalogRepository
}

func NewCatalogService(repo *CatalogRepository) *CatalogService {
	return &CatalogService{repo: repo}
}

func (s *CatalogService) GetParentCategories(ctx context.Context) ([]Category, error) {
	return s.repo.GetParentCategories(ctx)
}
