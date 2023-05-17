package svc

import (
	"github.com/iot-platform/admin/internal/config"
	"github.com/iot-platform/models"
	user2 "github.com/iot-platform/user/rpc/types/user"
	"github.com/iot-platform/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RpcUser  user.User
	AuthUser *user2.UserAuthReply
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: user.NewUser(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
