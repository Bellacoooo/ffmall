package service

import (
	"context"
	"ffmall/app/user/biz/dal/model"
	"ffmall/app/user/biz/dal/mysql"
	user "ffmall/rpc_gen/kitex_gen/user"
	"fmt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	user1 := model.User{}
	// Finish your business logic.
	result := mysql.DB.Where("email = ?", req.Email).First(&user1)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	if result.Error != nil {
		err = fmt.Errorf("账号找不到: %v", result.Error)
	} else {
		if user1.Password == req.Password {
			// 登陆成功，返回next页面
			resp = &user.LoginResp{
				UserId: int32(user1.ID),
			}
		} else {
			err = fmt.Errorf("密码错误")
		}
	}
	return
}
