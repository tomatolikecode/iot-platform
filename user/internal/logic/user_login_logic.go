package logic

import (
	"context"
	"errors"

	"github.com/iot-platform/helper"
	"github.com/iot-platform/models"
	"github.com/iot-platform/user/internal/svc"
	"github.com/iot-platform/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UserLoginReply)
	ub := new(models.UserBasic)
	err = l.svcCtx.DB.Where("name = ? AND password = ? ", req.UserName, req.Password).First(ub).Error
	if err != nil {
		logx.Error("[DB ERROR] :", err)
		err = errors.New("用户名或密码错误")
		return
	}
	token, err := helper.GenerateToken(ub.ID, ub.Identity, ub.Name, 3600*24*30)
	if err != nil {
		logx.Error("[TOKEN ERROR]:", err)
		err = errors.New("用户名或密码错误")
		return
	}

	resp.Token = token
	return
}
