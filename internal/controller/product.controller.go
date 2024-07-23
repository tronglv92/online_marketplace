package controler

import (
	"github.com/gin-gonic/gin"
	"github.com/online_marketplace/helper/server/http/response"
	"github.com/online_marketplace/internal/registry"
	"github.com/online_marketplace/internal/service"
	"github.com/online_marketplace/internal/types/request"
)

type ProductController struct {
	reg        *registry.ServiceContext
	productSvc service.ProductService
}

func NewProductController(reg *registry.ServiceContext) *ProductController {
	return &ProductController{
		reg:        reg,
		productSvc: service.NewProductService(reg),
	}
}

func (uc *ProductController) Create() func(*gin.Context) {
	return func(c *gin.Context) {
		var req request.ProductReq
		ctx := c.Request.Context()

		if err := c.ShouldBind(&req); err != nil {
			panic(err)
		}

		if err := req.Validate(ctx); err != nil {
			panic(err)

		}

		resp, err := uc.productSvc.Create(ctx, req)
		if err != nil {
			panic(err)
		}

		response.SuccessResponse(c, resp, nil)
	}

}

func (uc *ProductController) List() func(*gin.Context) {
	return func(c *gin.Context) {
		var req request.ProductListReq
		ctx := c.Request.Context()

		if err := c.ShouldBind(&req); err != nil {
			panic(err)
		}

		resp, paging, err := uc.productSvc.List(ctx, req)
		if err != nil {
			panic(err)
		}

		response.SuccessResponse(c, resp, paging)
	}

}
