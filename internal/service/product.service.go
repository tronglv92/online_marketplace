package service

import (
	"context"

	"github.com/online_marketplace/helper/errors"
	"github.com/online_marketplace/helper/model"

	"github.com/online_marketplace/internal/registry"
	"github.com/online_marketplace/internal/repository"
	"github.com/online_marketplace/internal/types/entity"
	"github.com/online_marketplace/internal/types/request"
	"github.com/online_marketplace/internal/types/response"
)

type ProductService interface {
	Create(ctx context.Context, input request.ProductReq) (*response.ProductResponse, error)
	List(ctx context.Context, input request.ProductListReq) ([]*response.ProductResponse, *model.Pagination, error)
}

type productSvcImpl struct {
	reg         *registry.ServiceContext
	productRepo repository.ProductRepository
}

func NewProductService(reg *registry.ServiceContext) ProductService {
	return &productSvcImpl{
		reg:         reg,
		productRepo: reg.ProductRepo,
	}
}

func (s *productSvcImpl) Create(ctx context.Context, input request.ProductReq) (*response.ProductResponse, error) {

	productEntity := entity.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Seller:      input.Seller,
	}
	product, err := s.productRepo.CreateWithReturn(ctx, &productEntity)
	if err != nil {
		return nil, errors.InternalServer(err)
	}

	return response.ProductMapToResponse(product), nil
}

func (s *productSvcImpl) List(ctx context.Context, input request.ProductListReq) ([]*response.ProductResponse, *model.Pagination, error) {
	products, pagination, err := s.productRepo.FindWithPagination(ctx, input.Limit, input.Page, s.productRepo.WithOrder(input.SortBy, input.SortOrder))
	if err != nil {
		return nil, nil, err
	}

	if pagination.TotalRecords == 0 {
		return nil, nil, errors.DataNotFound()
	}
	return response.ProductMapToResponses(products), pagination, nil
}
