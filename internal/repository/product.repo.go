package repository

import (
	"context"

	"github.com/online_marketplace/helper/database"
	baseRepo "github.com/online_marketplace/helper/database/sqldata/repository"
	"github.com/online_marketplace/helper/model"
	"github.com/online_marketplace/internal/types/entity"
)

type ProductRepository interface {
	FindWithPagination(ctx context.Context, limit int, page int, opts ...baseRepo.QueryOpt) ([]*entity.Product, *model.Pagination, error)
	CreateWithReturn(ctx context.Context, entity *entity.Product) (*entity.Product, error)
	WithOrder(sortBy string, sortOrder string, fields ...string) baseRepo.QueryOpt
}

type productRepo struct {
	baseRepo.BaseRepository[entity.Product]
}

func NewProductRepository(db database.Database) ProductRepository {
	return &productRepo{
		baseRepo.NewBaseRepository[entity.Product](db),
	}
}
