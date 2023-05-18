package logic

import (
	"context"

	"github.com/iot-platform/admin/internal/svc"
	"github.com/iot-platform/admin/internal/types"
	"github.com/iot-platform/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceModifyLogic {
	return &DeviceModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceModifyLogic) DeviceModify(req *types.DeviceModifyRequest) (resp *types.DeviceModifyReply, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.DB.Debug().Where("identity = ?", req.Identity).Updates(&models.DeviceBasic{
		ProductIdentity: req.ProductIdentity,
		Name:            req.Name,
	}).Error

	if err != nil {
		logx.Error("[DB ERROR]: ", err.Error())
	}
	return
}
