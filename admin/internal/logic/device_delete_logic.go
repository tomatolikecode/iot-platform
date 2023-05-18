package logic

import (
	"context"

	"github.com/iot-platform/admin/internal/svc"
	"github.com/iot-platform/admin/internal/types"
	"github.com/iot-platform/models"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceDeleteLogic {
	return &DeviceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceDeleteLogic) DeviceDelete(req *types.DeviceDeleteRequest) (resp *types.DeviceDeleteReply, err error) {
	// todo: add your logic here and delete this line
	deviceBase := new(models.DeviceBasic)
	err = l.svcCtx.DB.Debug().
		Model(new(models.DeviceBasic)).
		Select("key").
		Where("identity = ?", req.Identity).
		Find(deviceBase).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err.Error())
		return
	}

	err = l.svcCtx.DB.Debug().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("identity = ?", req.Identity).Delete(new(models.DeviceBasic)).Error; err != nil {
			return nil
		}
		return nil
	})

	return
}
