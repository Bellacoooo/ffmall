package home

import (
	"context"

	"ffmall/app/frontend/biz/utils"
	home "ffmall/app/frontend/hertz_gen/frontend/home"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &home.Empty{} //声明一个空的结构体，用于返回数据
	// resp, err = service.NewHomeService(ctx, c).Run(&req)
	// if err != nil {
	// 	utils.SendErrResponse(ctx, c, consts.StatusOK, err)
	// 	return
	// }
	c.Redirect(consts.StatusFound, []byte("/ping"))
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp1)
}
