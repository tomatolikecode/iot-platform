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

type ProductCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCreateLogic {
	return &ProductCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCreateLogic) ProductCreate(req *types.ProductCreateRequest) (resp *types.ProductCreateReply, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.DB.Debug().Transaction(func(tx *gorm.DB) error {
		product := models.ProductBasic{
			Identity: uuid.New().String(),
			Name:     req.Name,
			Desc:     req.Desc,
			Key:      uuid.New().String(),
		}

		if err := tx.Create(&product).Error; err != nil {
			return err
		}
		return nil
	})
	return
}
