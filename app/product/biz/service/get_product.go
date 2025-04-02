package service

import (
	"context"
	"ffmall/app/product/biz/dal/model"
	"ffmall/app/product/biz/dal/mysql"
	product "ffmall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	product1 := model.Product{}
	result := mysql.DB.Where("id = ?", req.Id).First(&product1)
	if result.Error != nil {
		return nil, result.Error
	}
	// int* resp;
	// int a = 0;
	// resp = &a;
	resp = &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(product1.ID),
			Name:        product1.Name,
			Picture:     product1.Picture,
			Price:       float32(product1.Price),
			Description: product1.Description,
		},
	}

	return
}
