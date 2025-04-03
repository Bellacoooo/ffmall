package service

import (
	"context"
	"ffmall/app/product/biz/dal/model"
	"ffmall/app/product/biz/dal/mysql"
	product "ffmall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	products := []model.Product{}
	mysql.DB.Where("name LIKE ?", "%"+req.Query+"%").Find(&products)
	if err != nil {
		return nil, err
	}

	resp = &product.SearchProductsResp{}
	for _, p := range products {
		resp.Results = append(resp.Results, &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Picture:     p.Picture,
			Price:       float32(p.Price),
			Description: p.Description,
		})
	}

	return
}
