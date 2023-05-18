package logic

import (
	"context"

	"github.com/iot-platform/device/internal/svc"
	"github.com/iot-platform/device/types/device"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *device.SendMessageRequest) (*device.SendMessageReply, error) {
	return &device.SendMessageReply{}, nil
}
