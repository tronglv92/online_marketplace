package request

import (
	"context"

	validation "github.com/itgelo/ozzo-validation"
)

type ProductReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Seller      string  `json:"seller"`
}

func (req ProductReq) Validate(ctx context.Context) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Price, validation.Required),
		validation.Field(&req.Seller, validation.Required),
	)
}

type ProductListReq struct {
	PaginationReq
	SortOrderReq
}
