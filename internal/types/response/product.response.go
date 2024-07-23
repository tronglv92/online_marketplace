package response

import "github.com/online_marketplace/internal/types/entity"

type ProductResponse struct {
	Id          int32   `json:"id"`
	UId         string  `json:"uid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Seller      string  `json:"seller"`
}

func ProductMapToResponse(data *entity.Product) *ProductResponse {
	return &ProductResponse{
		Id:          data.Id,
		UId:         data.UId,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Seller:      data.Seller,
	}
}

func ProductMapToResponses(items []*entity.Product) []*ProductResponse {
	var results []*ProductResponse
	for _, val := range items {
		results = append(results, ProductMapToResponse(val))
	}
	return results
}
