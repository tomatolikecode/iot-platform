package logic

import (
	"context"

	"github.com/google/uuid"
	"github.com/iot-platform/admin/internal/svc"
	"github.com/iot-platform/admin/internal/types"
	"github.com/iot-platform/models"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceCreateLogic {
	return &DeviceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceCreateLogic) DeviceCreate(req *types.DeviceCreateRequest) (resp *types.DeviceCreateReply, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.DB.Debug().Transaction(func(tx *gorm.DB) error {
		device := models.DeviceBasic{
			Identity:        uuid.New().String(),
			Name:            req.Name,
			ProductIdentity: req.ProductIdentity,
			Key:             uuid.New().String(),
			Secret:          uuid.New().String(),
		}

		if err := tx.Create(&device).Error; err != nil {
			return err
		}
		return nil
	})
	return
}
