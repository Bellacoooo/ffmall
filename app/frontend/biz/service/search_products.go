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

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp *map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	c, err := productservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	search_req := rpc_product.SearchProductsReq{
		Query: req.Q,
	}
	search_resp, err := c.SearchProducts(context.Background(), &search_req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	// for(int i = 0; i < Results.size(); i++) { result = Results[i];}
	var items []map[string]any
	for _, result := range search_resp.Results {
		search_resp_map := map[string]any{
			"ID":          result.Id,
			"Description": result.Description,
			"Name":        result.Name,
			"Picture":     result.Picture,
			"Price":       result.Price,
		}
		// items.push_back(search_resp_map)
		items = append(items, search_resp_map)
	}
	resp = &map[string]any{
		"items": items,
	}
	return
}
