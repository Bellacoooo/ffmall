package service

import (
	"context"
	"errors"
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

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.SignUpReq) (resp string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	if req.Password != req.PasswordConfirm {
		err = errors.New("Passwords do not match")
		return
	}
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	c, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	register_req := user.RegisterReq{
		Email:    req.Email,
		Password: req.Password,
	}
	register_resp, err := c.Register(context.Background(), &register_req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_ID", register_resp.UserId)
	_ = session.Save()
	resp = "注册成功"
	return
}
