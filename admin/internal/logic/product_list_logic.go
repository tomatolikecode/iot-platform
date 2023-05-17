package logic

import (
	"context"

	"github.com/iot-platform/admin/internal/svc"
	"github.com/iot-platform/admin/internal/types"
	"github.com/iot-platform/helper"
	"github.com/iot-platform/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ProductListRequest) (resp *types.ProductListReply, err error) {
	// todo: add your logic here and delete this line
	logx.Info(req)
	list := make([]*types.ProductListBasic, 0)
	req.Size = helper.If(req.Size == 0, 20, req.Size).(int)
	req.Page = helper.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)
	var count int64
	resp = new(types.ProductListReply)

	err = models.ProductList(req.Name).Count(&count).Offset(req.Page).Limit(req.Size).Find(&list).Error
	if err != nil {
		logx.Error("[DB ERROR]: ", err.Error())
		return
	}

	for _, v := range list {
		v.CreatedAt = helper.RFC3339ToNormalTime(v.CreatedAt)
	}

	resp.Count = count
	resp.List = list
	return
}
