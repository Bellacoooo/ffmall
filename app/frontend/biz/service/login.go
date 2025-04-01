package service

import (
	"context"
	"log"
	"time"

	auth "ffmall/app/frontend/hertz_gen/frontend/auth"
	"ffmall/rpc_gen/kitex_gen/user"
	"ffmall/rpc_gen/kitex_gen/user/userservice"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/hertz-contrib/sessions"
	consul "github.com/kitex-contrib/registry-consul"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// 调用User服务的Login方法
	// rpc_resp := user.Login(req)
	// user_id := rpc_resp.UserID
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	c, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	login_req := user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	}
	login_resp, err := c.Login(context.Background(), &login_req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_ID", login_resp.UserId)
	_ = session.Save()

	resp = req.Next

	return resp, err
}
