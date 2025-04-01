package service

import (
	"context"
	"ffmall/app/user/biz/dal/model"
	"ffmall/app/user/biz/dal/mysql"
	user "ffmall/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.

	user1 := &model.User{
		Email:    &req.Email,
		Password: req.Password,
	}
	result := mysql.DB.Create(user1)
	if result.Error != nil {
		err = result.Error
		return
	}
	resp = &user.RegisterResp{
		UserId: int32(user1.ID),
	}

	return
}
