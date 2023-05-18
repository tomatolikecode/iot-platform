package logic

import (
	"context"

	"github.com/iot-platform/admin/internal/svc"
	"github.com/iot-platform/admin/internal/types"
	"github.com/iot-platform/models"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDeleteLogic) ProductDelete(req *types.ProductDeleteRequest) (resp *types.ProductDeleteReply, err error) {
	// todo: add your logic here and delete this line
	productBasic := new(models.ProductBasic)
	err = l.svcCtx.DB.Debug().
		Model(new(models.ProductBasic)).
		Select("key").
		Where("identity = ?", req.Identity).
		Find(productBasic).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err.Error())
		return
	}

	err = l.svcCtx.DB.Debug().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("identity = ?", req.Identity).Delete(new(models.ProductBasic)).Error; err != nil {
			return nil
		}
		return nil
	})

	return
}
