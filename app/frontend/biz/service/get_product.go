package service

import (
	"context"
	"log"
	"time"

	product "ffmall/app/frontend/hertz_gen/frontend/product"
	rpc_product "ffmall/rpc_gen/kitex_gen/product"
	"ffmall/rpc_gen/kitex_gen/product/productservice"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	consul "github.com/kitex-contrib/registry-consul"
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

	// resp := map[string]any{
	// 	"item" :map[string]any{
	// 		"Picture": xxx,
	// 		"Name": xxx,
	// 		.....
	// 	}
	// }
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	c, err := productservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	product_req := rpc_product.GetProductReq{
		Id: int64(req.Id),
	}
	product_resp, err := c.GetProduct(context.Background(), &product_req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	product_resp_map := map[string]any{
		"ID":          product_resp.Product.Id,
		"Description": product_resp.Product.Description,
		"Name":        product_resp.Product.Name,
		"Picture":     product_resp.Product.Picture,
		"Price":       product_resp.Product.Price,
	}
	resp = &map[string]any{
		"item": product_resp_map,
	}
	return
}
