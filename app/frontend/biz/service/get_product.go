package service

import (
	"context"

	"ffmall/app/frontend/biz/dal/model"
	"ffmall/app/frontend/biz/dal/mysql"
	product "ffmall/app/frontend/hertz_gen/frontend/product"

	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp *map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	product1 := model.Product{}
	result := mysql.DB.Where("id = ?", req.Id).First(&product1)
	if result.Error != nil {
		return nil, result.Error
	}
	resp = &map[string]any{}
	p := map[string]any{
		"Name":    product1.Name,
		"Price":   product1.Price,
		"Picture": product1.Picture,
		"Id":      product1.ID,
	}
	(*resp)["item"] = p
	// resp := map[string]any{
	// 	"item" :map[string]any{
	// 		"Picture": xxx,
	// 		"Name": xxx,
	// 		.....
	// 	}
	// }
	return
}
